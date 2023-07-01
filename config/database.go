package config

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:admin@/go_products?parseTime=true")
	if err != nil {
		panic(err)
	}
	log.Println("Database Connected")
	DB = db
}
