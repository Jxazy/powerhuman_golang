package test

import (
	"context"
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
	"jxazy/powerhuman_golang/model/domain"
	"jxazy/powerhuman_golang/repository"
	"jxazy/powerhuman_golang/service"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func truncateEmployee(db *sql.DB) {
	db.Exec("TRUNCATE employees")
}

func TestCreateEmployeeSuccess(t *testing.T) {
	db := setupTestDB()
	truncateEmployee(db)
	router := SetupRouter(db)
	requestBody := strings.NewReader(`{"name" : "Golang Unit Test 1", "email" : "golang1@gmail.com", "gender": "MALE", "age": 20, "phone": "393728", "team_id": 34, "role_id": 52}`)
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

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Golang Unit Test 1", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "golang1@gmail.com", responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, "MALE", responseBody["data"].(map[string]interface{})["gender"])
	assert.Equal(t, 20, int(responseBody["data"].(map[string]interface{})["age"].(float64)))
	assert.Equal(t, "393728", responseBody["data"].(map[string]interface{})["phone"])
	assert.Equal(t, 34, int(responseBody["data"].(map[string]interface{})["team_id"].(float64)))
	assert.Equal(t, 52, int(responseBody["data"].(map[string]interface{})["role_id"].(float64)))
}

func TestCreateEmployeeFailed(t *testing.T) {
	db := setupTestDB()
	truncateEmployee(db)
	router := SetupRouter(db)
	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/employees", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateEmployeeSuccess(t *testing.T) {
	db := setupTestDB()
	truncateEmployee(db)

	tx, _ := db.Begin()
	employeeRepository := repository.NewEmployeeRepositoryImpl()
	employee := employeeRepository.Save(context.Background(), tx, domain.Employee{
		Name:   "Golang Unit Test Update",
		Email:  "golangupdate@gmail.com",
		Gender: "MALE",
		Age:    20,
		Phone:  "393728",
		TeamId: 34,
		RoleId: 52,
	})

	tx.Commit()

	router := SetupRouter(db)
	requestBody := strings.NewReader(`{"name" : "Golang Unit Test Update", "email" : "golangupdate@gmail.com", "gender": "MALE", "age": 20, "phone": "393728", "team_id": 34, "role_id": 52}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/employees/"+strconv.Itoa(employee.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, employee.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Golang Unit Test Update", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "golangupdate@gmail.com", responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, "MALE", responseBody["data"].(map[string]interface{})["gender"])
	assert.Equal(t, 20, int(responseBody["data"].(map[string]interface{})["age"].(float64)))
	assert.Equal(t, "393728", responseBody["data"].(map[string]interface{})["phone"])
	assert.Equal(t, 34, int(responseBody["data"].(map[string]interface{})["team_id"].(float64)))
	assert.Equal(t, 52, int(responseBody["data"].(map[string]interface{})["role_id"].(float64)))
}

func TestUpdateEmployeeFailed(t *testing.T) {
	db := setupTestDB()
	truncateEmployee(db)

	tx, _ := db.Begin()
	employeeRepository := repository.NewEmployeeRepositoryImpl()
	employee := employeeRepository.Save(context.Background(), tx, domain.Employee{
		Name:   "Golang Unit Test Update",
		Email:  "golangupdate@gmail.com",
		Gender: "MALE",
		Age:    20,
		Phone:  "393728",
		TeamId: 34,
		RoleId: 52,
	})

	tx.Commit()

	router := SetupRouter(db)
	requestBody := strings.NewReader(`{"name" : "", "email" : "golang1@gmail.com", "gender": "MALE", "age": 20, "phone": "393728", "team_id": 34, "role_id": 52}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/employees/"+strconv.Itoa(employee.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
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
