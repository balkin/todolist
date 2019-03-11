package db

import (
	. "github.com/balkin/todolist/todo"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var Db *pg.DB

func ConnectToDatabase() {
	Db = pg.Connect(&pg.Options{User: "root", Password: "Passw0rd", Addr: "postgres:5432"})
	if err := Db.CreateTable((*TodoItem)(nil), &orm.CreateTableOptions{IfNotExists: true, FKConstraints: true}); err != nil {
		panic(err)
	}
}

func ConnectToTestDatabase() {
	Db = pg.Connect(&pg.Options{User: "root", Password: "Passw0rd"})
	if err := Db.CreateTable((*TodoItem)(nil), &orm.CreateTableOptions{Temp: true, IfNotExists: true, FKConstraints: true}); err != nil {
		panic(err)
	}
}

func DisconnectDatabase() {
	if Db != nil {
		_ = Db.Close()
		Db = nil
	}
}
