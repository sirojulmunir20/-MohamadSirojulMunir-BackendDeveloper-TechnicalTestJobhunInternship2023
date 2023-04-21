package config

import (
	"database/sql"
	"fmt"
	"log"
  _ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = ""
	host     = "localhost"
	port     = "3306"
	dbname   = "jobhun_test"
)

func NewDB() (*sql.DB, error) {
	// create the data source name
	dsn:= fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	// open  database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open connection database: %v", err)
		return nil,err
	}

	// check if the connection is valid
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Connected to database!")
	return db, nil
}