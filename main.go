package main

import (
	"fmt"
	"jobhun-api/config"
	"jobhun-api/helper"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Start server")
	
	db, err := config.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	
	defer db.Close()
	
	routes := httprouter.New()

	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
	
}