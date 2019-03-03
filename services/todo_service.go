package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type TodoItem struct {
	gorm.Model
	SubItems []*TodoItem `gorm:"many2many:todo_subitems;association_jointable_foreignkey:todo_subitem_id"`
}

var Db *gorm.DB

func ConnectToDatabase() error {
	var err error
	Db, err = gorm.Open("postgres", "host=localhost user=root dbname=root password=Passw0rd sslmode=disable")
	if err != nil {
		return err
	}
	Db.AutoMigrate(&TodoItem{})
	return err
}

func DisconnectDatabase() error {
	if Db != nil {
		err := Db.Close()
		Db = nil
		return err
	}
	return nil
}
