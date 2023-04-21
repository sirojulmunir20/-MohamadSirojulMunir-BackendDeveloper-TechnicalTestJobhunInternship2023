package router

import (
	"fmt"
	"jobhun-api/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(studentController *controller.StudentController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	router.GET("/api/student", studentController.FindAll)
	router.GET("/api/student/:studentId", studentController.FindById)
	router.POST("/api/student", studentController.Create)
	router.PATCH("/api/student/:studentId", studentController.Update)
	router.DELETE("/api/student/:studentId", studentController.Delete)

	return router
}