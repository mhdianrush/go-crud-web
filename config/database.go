package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()
var DB *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		logger.Printf("failed loaded env file %s", err.Error())
	}
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"),
	)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		logger.Printf("failed connect to database %s", err.Error())
	}

	logger.Println("Database Connected")

	DB = db
}
