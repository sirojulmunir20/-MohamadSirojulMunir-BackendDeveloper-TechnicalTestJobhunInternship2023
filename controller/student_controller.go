package controller


import (
	"jobhun-api/data/request"
	"jobhun-api/data/response"
	"jobhun-api/helper"
	"jobhun-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type StudentController struct {
	StudentService service.StudentService
}

func NewStudentController(studentService service.StudentService) *StudentController {
	return &StudentController{StudentService: studentService}
}

func (controller *StudentController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	studentCreateRequest := request.StudentCreateRequest{}
	helper.ReadRequestBody(requests, &studentCreateRequest)

	controller.StudentService.Create(requests.Context(), studentCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *StudentController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	studentUpdateRequest := request.StudentUpdateRequest{}
	helper.ReadRequestBody(requests, &studentUpdateRequest)

	studentId := params.ByName("studentId")
	id, err := strconv.Atoi(studentId)
	helper.PanicIfError(err)
	studentUpdateRequest.Id = id

	controller.StudentService.Update(requests.Context(), studentUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *StudentController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	studentId := params.ByName("studentId")
	id, err := strconv.Atoi(studentId)
	helper.PanicIfError(err)

	controller.StudentService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *StudentController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.StudentService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *StudentController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	studentId := params.ByName("studentId")
	id, err := strconv.Atoi(studentId)
	helper.PanicIfError(err)

	result := controller.StudentService.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)

}