package server

import (
	"github.com/goraft/raft"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type UpdateTableCommand struct {
	T string
	N string
	C int
	W string
}

func NewUpdateTableCommand(txid string, name string, friendCount int, favoriteWord string) *UpdateTableCommand {
	return &UpdateTableCommand{
		T: txid,
		N: name,
		C: friendCount,
		W: favoriteWord,
	}
}

func (this *UpdateTableCommand) CommandName() string {
	return "update"
}

func (this *UpdateTableCommand) Apply(server raft.Server) (interface{}, error) {
	db := server.Context().(*Store)
	return db.Update(this.T, this.N, this.C, this.W), nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type CreateTableCommand struct {
	T string
}

func NewCreateTableCommand(txid string) *CreateTableCommand {
	return &CreateTableCommand{
		T: txid,
	}
}

func (this *CreateTableCommand) CommandName() string {
	return "create"
}

func (this *CreateTableCommand) Apply(server raft.Server) (interface{}, error) {
	db := server.Context().(*Store)
	return db.Create(this.T), nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
