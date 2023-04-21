package main

import (
	"fmt"
	"jobhun-api/config"
	"jobhun-api/helper"
	"jobhun-api/router"
	"jobhun-api/repository"
	"jobhun-api/service"
	"jobhun-api/controller"
	
	"net/http"

	_"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Start server")
	
	db := config.DatabaseConnection()
	defer db.Close()
	
	// repository
	studentRepository := repository.NewStudentRepository(db)

	// service
	studentService := service.NewStudentServiceImpl(studentRepository)

	// controller
	bookController := controller.NewStudentController(studentService)

	// router
	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
	
}