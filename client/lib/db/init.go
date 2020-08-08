package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// InitDB DataBase Connection
func InitDB() {
	var err error
	db, err = gorm.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalln(err)
		return
	}
	var models = []interface{}{&UserInfo{}, &ServerInfo{}}
	db.AutoMigrate(models...)
}

func DB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
