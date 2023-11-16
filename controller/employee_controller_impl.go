package controller

import (
	"github.com/julienschmidt/httprouter"
	"jxazy/powerhuman_golang/helper"
	"jxazy/powerhuman_golang/model/web"
	"jxazy/powerhuman_golang/service"
	"net/http"
	"strconv"
)

type EmployeeControllerImpl struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeController(employeeService service.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{
		EmployeeService: employeeService,
	}
}

func (controller EmployeeControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	employeeCreateRequest := web.EmployeeCreateRequest{}
	helper.ReadFromRequestBody(r, &employeeCreateRequest)

	employeeResponse := controller.EmployeeService.Create(r.Context(), employeeCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller EmployeeControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	employeeUpdateRequest := web.EmployeeUpdateRequest{}
	helper.ReadFromRequestBody(r, &employeeUpdateRequest)

	employeeId := params.ByName("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.CheckError(err)

	employeeUpdateRequest.Id = id

	employeeResponse := controller.EmployeeService.Update(r.Context(), employeeUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller EmployeeControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	employeeId := params.ByName("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.CheckError(err)

	controller.EmployeeService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller EmployeeControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	employeeId := params.ByName("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.CheckError(err)

	employeeResponse := controller.EmployeeService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller EmployeeControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	employeeResponse := controller.EmployeeService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}
