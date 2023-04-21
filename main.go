package main

import (
	"jobhun-api/config"
	"log"
)

func main() {
	db, err := config.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer db.Close()
}