package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"jxazy/powerhuman_golang/app"
	"jxazy/powerhuman_golang/controller"
	"jxazy/powerhuman_golang/helper"
	"jxazy/powerhuman_golang/middleware"
	"jxazy/powerhuman_golang/repository"
	"jxazy/powerhuman_golang/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	employeeRepository := repository.NewEmployeeRepositoryImpl()
	employeeService := service.NewEmployeeServiceImpl(employeeRepository, db, validate)
	employeeController := controller.NewEmployeeController(employeeService)

	router := app.NewRouter(employeeController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.CheckError(err)
}
