package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"jxazy/powerhuman_golang/exception"
	"jxazy/powerhuman_golang/helper"
	"jxazy/powerhuman_golang/model/domain"
	"jxazy/powerhuman_golang/model/web"
	"jxazy/powerhuman_golang/repository"
)

type EmployeeServiceImpl struct {
	EmployeeRepository repository.EmployeeRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

func NewEmployeeServiceImpl(employeeRepository repository.EmployeeRepository, DB *sql.DB, validator *validator.Validate) EmployeeService {
	return &EmployeeServiceImpl{
		EmployeeRepository: employeeRepository,
		DB:                 DB,
		Validator:          validator,
	}
}

func (service EmployeeServiceImpl) Create(ctx context.Context, request web.EmployeeCreateRequest) web.EmployeeResponse {
	err := service.Validator.Struct(request)
	helper.CheckError(err)

	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	employee := domain.Employee{
		Name:   request.Name,
		Email:  request.Email,
		Gender: request.Gender,
		Age:    request.Age,
		Phone:  request.Phone,
		TeamId: request.TeamId,
		RoleId: request.RoleId,
	}

	employee = service.EmployeeRepository.Save(ctx, tx, employee)

	return helper.ToEmployeeResponse(employee)
}

func (service EmployeeServiceImpl) Update(ctx context.Context, request web.EmployeeUpdateRequest) web.EmployeeResponse {
	err := service.Validator.Struct(request)
	helper.CheckError(err)

	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := service.EmployeeRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	employee.Name = request.Name
	employee.Email = request.Email
	employee.Gender = request.Gender
	employee.Age = request.Age
	employee.Phone = request.Phone
	employee.TeamId = request.TeamId
	employee.RoleId = request.RoleId

	service.EmployeeRepository.Update(ctx, tx, employee)

	return helper.ToEmployeeResponse(employee)
}

func (service EmployeeServiceImpl) Delete(ctx context.Context, employeeId int) {
	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.EmployeeRepository.FindById(ctx, tx, employeeId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	service.EmployeeRepository.Delete(ctx, tx, category)
}

func (service EmployeeServiceImpl) FindById(ctx context.Context, employeeId int) web.EmployeeResponse {
	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.EmployeeRepository.FindById(ctx, tx, employeeId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	return helper.ToEmployeeResponse(category)
}

func (service EmployeeServiceImpl) FindAll(ctx context.Context) []web.EmployeeResponse {
	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.EmployeeRepository.FindAll(ctx, tx)

	return helper.ToEmployeeResponses(categories)
}
