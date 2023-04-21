package config

import (
	"database/sql"
	"fmt"
	"jobhun-api/helper"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = ""
	host     = "localhost"
	port     = "3306"
	dbname   = "jobhun_test"
)

func DatabaseConnection() *sql.DB {
	dsn:= fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	fmt.Println("Connected to database!")

	return db
}