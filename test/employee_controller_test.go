package test

import (
	"database/sql"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"jxazy/powerhuman_golang/app"
	"jxazy/powerhuman_golang/controller"
	"jxazy/powerhuman_golang/helper"
	"jxazy/powerhuman_golang/middleware"
	"jxazy/powerhuman_golang/repository"
	"jxazy/powerhuman_golang/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/powerhuman-backend-test")
	helper.CheckError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func SetupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	employeeRepository := repository.NewEmployeeRepositoryImpl()
	employeeService := service.NewEmployeeServiceImpl(employeeRepository, db, validate)
	employeeController := controller.NewEmployeeController(employeeService)

	router := app.NewRouter(employeeController)

	return middleware.NewAuthMiddleware(router)
}

func truncateEmploye(db *sql.DB) {
	db.Exec("TRUNCATE employees")
}

func TestCreateEmployeeSuccess(t *testing.T) {
	db := setupTestDB()
	router := SetupRouter(db)
	requestBody := strings.NewReader(`{"name" : "Gadget", "email" : "asuhas@gmail.com", }`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/employees", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Golang Unit Test 1", responseBody["name"])

}

func TestCreateEmployeeFailed(t *testing.T) {

}

func TestUpdateEmployeeSuccess(t *testing.T) {

}

func TestUpdateEmployeeFailed(t *testing.T) {

}

func TestGetEmployeeSuccess(t *testing.T) {

}

func TestGetEmployeeFailed(t *testing.T) {

}

func TestDeleteEmployeeSuccess(t *testing.T) {

}

func TestDeleteEmployeeFailed(t *testing.T) {

}

func TestGetListEmployeesSucces(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {

}
