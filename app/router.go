package app

import (
	"github.com/julienschmidt/httprouter"
	"jxazy/powerhuman_golang/controller"
	"jxazy/powerhuman_golang/exception"
)

func NewRouter(employeeController controller.EmployeeController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/employees", employeeController.FindAll)
	router.GET("/api/employees/:employeeId", employeeController.FindById)
	router.POST("/api/employees", employeeController.Create)
	router.PUT("/api/employees/:employeeId", employeeController.Update)
	router.DELETE("/api/employees/:employeeId", employeeController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
