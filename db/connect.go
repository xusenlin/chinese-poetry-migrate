package db


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Conn *gorm.DB

func InitConn() {
	var err error

	Conn, err = gorm.Open("sqlite3", "./database.db")

	if err != nil {
		panic(err)
	}
	Conn.SingularTable(true)
}