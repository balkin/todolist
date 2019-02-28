package services

import (
	"github.com/go-pg/pg"
	"log"
)

var db *pg.DB

func ConnectToDatabase() {
	db = pg.Connect(&pg.Options{
		User: "postgres",
	})
}

func DisconnectDatabase() {
	if db == nil {
		return
	}
	if err := db.Close(); err != nil {
		log.Println("Failed to disconnect database", err)
	}
}
