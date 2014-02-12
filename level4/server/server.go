package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/goraft/raft"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"stripe-ctf.com/sqlcluster/transport"
	"sync"
	"time"
)

var delay = 10 * time.Millisecond

type Server struct {
	name             string
	listen           string
	connectionString string
	router           *mux.Router
	raft             raft.Server
	http             *http.Server
	client           *transport.Client
	db               *Store
	txid             int
	mutex            sync.Mutex
}

func NewServer(listen string, directory string) (*Server, error) {
	raft.RegisterCommand(&CreateTableCommand{})
	raft.RegisterCommand(&UpdateTableCommand{})
	// raft.SetLogLevel(raft.Trace)

	connectionString, err := transport.Encode(listen)
	if err != nil {
		return nil, err
	}

	name := fmt.Sprintf("%d", rand.New(rand.NewSource(time.Now().UnixNano())).Int()%100)
	//log.Printf("name: %s", name)
	router := mux.NewRouter()
	db := NewStore()

	transporter := raft.NewHTTPTransporter("/r")
	transporter.DisableKeepAlives = true
	transporter.Transport.Dial = transport.UnixDialer

	raft, err := raft.NewServer(name, directory, transporter, nil, db, "")
	if err != nil {
		log.Fatal(err)
	}
	raft.SetElectionTimeout(150 * time.Millisecond)
	raft.SetHeartbeatTimeout(50 * time.Millisecond)
	//raft.SetElectionTimeout(5 * time.Second)
	//raft.SetHeartbeatTimeout(1 * time.Second)

	s := &Server{
		name:             name,
		listen:           listen,
		connectionString: connectionString,
		router:           router,
		raft:             raft,
		http:             &http.Server{Handler: router},
		client:           transport.NewClient(),
		db:               db,
		txid:             0,
	}

	transporter.Install(raft, s)

	return s, nil
}

func (this *Server) ListenAndServe(leader string) error {
	// Start Unix transport
	l, err := transport.Listen(this.listen)
	if err != nil {
		log.Fatal(err)
	}

	// Start the raft server
	//log.Printf("Starting raft server...")
	this.raft.Start()

	// If we're a follower, then join our leader
	if leader != "" {
		if err := this.Join(leader); err != nil {
			log.Fatal(err)
		}
	} else {
		//log.Printf("Initializing new cluster...")
		_, err := this.raft.Do(&raft.DefaultJoinCommand{
			Name:             this.raft.Name(),
			ConnectionString: this.connectionString,
		})
		if err != nil {
			log.Fatal(err)
		}
		//log.Printf("Done.")
	}
	go func() {
		for {
			leaderName := this.raft.Leader()
			if this.name == leaderName {
				//log.Printf("Leader is me!")
			} else {
				//log.Printf("Leader is %s", leaderName)
				if leader := this.raft.Peers()[leaderName]; leader != nil {
					//log.Printf("Leader url is %s", leader.ConnectionString)
				}
			}

			time.Sleep(1 * time.Second)
		}
	}()
	//log.Printf("Raft started.")

	//log.Printf("Starting http server on %s...", this.listen)
	this.router.HandleFunc("/sql", this.onSQL).Methods("POST")
	this.router.HandleFunc("/s", this.onParsedSQL).Methods("POST")
	this.router.HandleFunc("/j", this.onJoin).Methods("POST")
	//log.Printf("Done.")

	return this.http.Serve(l)
}

func (this *Server) GetLeaderConnectionString() (string, error) {
	if this.raft.State() == raft.Leader {
		return "", nil
	}

	name := this.raft.Leader()
	if leader := this.raft.Peers()[name]; leader != nil {
		return leader.ConnectionString, nil
	} else {
		return "", errors.New("Unable to determine leader")
	}
}

// This is a hack around Gorilla mux not providing the correct net/http
// HandleFunc() interface.
func (this *Server) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	this.router.HandleFunc(pattern, handler)
}

// Joins the leader of an existing cluster.
func (this *Server) Join(leader string) error {
	//log.Printf("Joining leader %s...", leader)
	command := &raft.DefaultJoinCommand{
		Name:             this.raft.Name(),
		ConnectionString: this.connectionString,
	}

	var b bytes.Buffer
	for {
		json.NewEncoder(&b).Encode(command)
		_, err := this.client.SafePost(leader, "/j", &b)
		if err != nil {
			//log.Printf("Error joining leader: %s", err)
			continue
		}

		return nil
	}
}

func (this *Server) onJoin(w http.ResponseWriter, req *http.Request) {
	//log.Printf("onJoin called")
	command := &raft.DefaultJoinCommand{}

	if err := json.NewDecoder(req.Body).Decode(&command); err != nil {
		//log.Printf("Error while decoding during onJoin: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := this.raft.Do(command); err != nil {
		//log.Printf("Error while executing raft command during onJoin: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Printf("Processed join command: %s", command)
}

//
// SQL methods
//

func (this *Server) onSQL(w http.ResponseWriter, req *http.Request) {
	query, err := ioutil.ReadAll(req.Body)
	if err != nil {
		//log.Printf("[name=%s] Couldn't read body: %s", this.name, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//log.Printf("[name=%s] Received query: %s", this.name, string(query))

	friendCountString, favoriteWord, name := this.db.ParseSQL(string(query))
	//log.Printf("[name=%s] name: %s, friendCount: %s, favoriteWord: %s", this.name, name, friendCountString, favoriteWord)

	var friendCount = 0
	if friendCountString != "" {
		num, err := strconv.Atoi(friendCountString)
		if err != nil {
			//log.Printf("[name=%s] Couldn't parse integer first occurrence: %s", this.name, friendCountString)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		friendCount = num
	}

	// Determine the transaction id that we may need to use for this query, worst case we burn a transaction id that we'll
	// never use because we're the leader.
	this.mutex.Lock()
	this.txid = this.txid + 1
	txid := fmt.Sprintf("%s:%d", this.name, this.txid)
	this.mutex.Unlock()

	// Compute the short request just in case we need it
	shortRequest := fmt.Sprintf("%s|%s|%s|%s", txid, friendCountString, favoriteWord, name)

	// Keep trying until we execute the query
	for {
		// Figure out which node is the leader because we can't just execute a Do command on a follower -- if we do then the
		// follower will just return a HTTP redirect to the leader which octopus won't follow.
		leader, err := this.GetLeaderConnectionString()
		if err != nil {
			////log.Printf("[name=%s] Unable to determine leader, retrying...", this.name)
			time.Sleep(delay)
			continue
		}

		if leader == "" {
			// I am the leader, I can execute a Do command
			var command raft.Command
			if friendCountString == "" && favoriteWord == "" && name == "" {
				command = NewCreateTableCommand(txid)
			} else {
				command = NewUpdateTableCommand(txid, name, friendCount, favoriteWord)
			}

			response, err := this.raft.Do(command)
			if err != nil {
				//log.Printf("Got an error while executing command: %s", err.Error())
				time.Sleep(delay)
				continue
			}

			// Return the response to the caller
			w.Write([]byte(response.(string)))
			break
		}

		if leader != "" {
			// We're not the leader, send the parsed request to the leader
			response, err := this.client.SafePost(leader, "/s", bytes.NewBuffer([]byte(shortRequest)))
			if err != nil {
				//log.Printf("[name=%s] Received error: %s, resending request for txid=%d", this.name, err, txid)
				time.Sleep(delay)
				continue
			}

			bs, _ := ioutil.ReadAll(response)
			w.Write(bs)
			break
		}
	}
}

func (this *Server) onParsedSQL(w http.ResponseWriter, req *http.Request) {
	query, err := ioutil.ReadAll(req.Body)
	if err != nil {
		//log.Printf("[name=%s] Couldn't read body: %s", this.name, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//log.Printf("[name=%s] Received parsed query: %#v", this.name, string(query))

	splits := strings.SplitN(string(query), "|", 4)
	txid := splits[0]
	friendCountString := splits[1]
	favoriteWord := splits[2]
	name := splits[3]

	// Parse the friendCount
	var friendCount = 0
	if friendCountString != "" {
		num, err := strconv.Atoi(friendCountString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		friendCount = num
	}

	for {
		// Figure out which node is the leader because we can't just execute a Do command on a follower -- if we do then the
		// follower will just return a HTTP redirect to the leader which octopus won't follow.
		leader, err := this.GetLeaderConnectionString()
		if err != nil {
			time.Sleep(delay)
			continue
		}

		if leader == "" {
			// I am the leader, I can execute a Do command
			var command raft.Command
			if friendCountString == "" && favoriteWord == "" && name == "" {
				command = NewCreateTableCommand(txid)
			} else {
				command = NewUpdateTableCommand(txid, name, friendCount, favoriteWord)
			}

			// Execute the command
			response, err := this.raft.Do(command)
			if err != nil {
				//log.Printf("Got an error while executing command: %s", err.Error())
				time.Sleep(delay)
				continue
			}

			// Return the response to the caller
			w.Write([]byte(response.(string)))
			break
		}

		if leader != "" {
			// We're not the leader, send the parsed request to the leader
			response, err := this.client.SafePost(leader, "/s", bytes.NewBuffer(query))
			if err != nil {
				//log.Printf("[name=%s] Received error: %s, resending request for txid=%d", this.name, err, txid)
				time.Sleep(delay)
				continue
			}

			bs, _ := ioutil.ReadAll(response)
			w.Write(bs)
			break
		}
	}
}
