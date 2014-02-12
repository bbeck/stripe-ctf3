package server

import (
	"fmt"
	"regexp"
	"sync"
)

type Row struct {
	friendCount  int
	requestCount int
	favoriteWord string
}

type Store struct {
	sequence     int
	rows         map[string]*Row
	createdTable bool
	cache        map[string]string
	mutex        sync.Mutex
}

var create = regexp.MustCompile("CREATE TABLE .*")
var update = regexp.MustCompile(".*UPDATE ctf3 SET friendCount=friendCount.([0-9]+).*favoriteWord=.([a-z]+). WHERE name=.([a-z]+).;.*")

func NewStore() *Store {
	rows := map[string]*Row{
		"siddarth":  &Row{friendCount: 0, requestCount: 0, favoriteWord: ""},
		"gdb":       &Row{friendCount: 0, requestCount: 0, favoriteWord: ""},
		"christian": &Row{friendCount: 0, requestCount: 0, favoriteWord: ""},
		"andy":      &Row{friendCount: 0, requestCount: 0, favoriteWord: ""},
		"carl":      &Row{friendCount: 0, requestCount: 0, favoriteWord: ""},
	}

	return &Store{
		sequence:     -1,
		rows:         rows,
		createdTable: false,
		cache:        make(map[string]string),
	}
}

func (this *Store) ParseSQL(sql string) (string, string, string) {
	if create.MatchString(sql) {
		return "", "", ""
	}

	if update.MatchString(sql) {
		matches := update.FindStringSubmatch(sql)
		friendCount := matches[1]
		favoriteWord := matches[2]
		name := matches[3]
		return friendCount, favoriteWord, name
	}

	panic(fmt.Sprintf("Unable to parse sql: %s", sql))
}

func (this *Store) Create(txid string) string {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// See if we have a cache hit
	last, contains := this.cache[txid]
	if contains {
		return last
	}

	// This is a new query, allocate the squence number and continue executing it
	this.sequence = this.sequence + 1
	//log.Printf("Allocated sequence number: %d", this.sequence)

	var response string
	if !this.createdTable {
		// First time, just return the empty string
		this.createdTable = true
		response = fmt.Sprintf("SequenceNumber: %d\n", this.sequence)
	} else {
		// We've done this already
		response = fmt.Sprintf("SequenceNumber: %d\nError: near line 1: table ctf3 already exists\nError: near line 6: column name is not unique\n", this.sequence)
	}

	this.cache[txid] = response
	return response
}

func (this *Store) Update(txid string, name string, friendCount int, favoriteWord string) string {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// See if we have a cache hit
	last, contains := this.cache[txid]
	if contains {
		return last
	}

	// This is a new query, allocate the squence number and continue executing it
	this.sequence = this.sequence + 1
	//log.Printf("Allocated sequence number=%d in response to update name:%s, friendCount:%d, favoriteWord:%s", this.sequence, name, friendCount, favoriteWord)

	row := this.rows[name]
	row.friendCount += friendCount
	row.requestCount += 1
	row.favoriteWord = favoriteWord

	s := this.rows["siddarth"]
	g := this.rows["gdb"]
	c := this.rows["christian"]
	a := this.rows["andy"]
	k := this.rows["carl"]

	response := fmt.Sprintf("SequenceNumber: %d\n"+
		"siddarth|%d|%d|%s\n"+
		"gdb|%d|%d|%s\n"+
		"christian|%d|%d|%s\n"+
		"andy|%d|%d|%s\n"+
		"carl|%d|%d|%s\n",
		this.sequence,
		s.friendCount, s.requestCount, s.favoriteWord,
		g.friendCount, g.requestCount, g.favoriteWord,
		c.friendCount, c.requestCount, c.favoriteWord,
		a.friendCount, a.requestCount, a.favoriteWord,
		k.friendCount, k.requestCount, k.favoriteWord,
	)

	this.cache[txid] = response
	return response
}
