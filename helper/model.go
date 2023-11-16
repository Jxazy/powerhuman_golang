package helper

import (
	"jxazy/powerhuman_golang/model/domain"
	"jxazy/powerhuman_golang/model/web"
)

func ToEmployeeResponse(employee domain.Employee) web.EmployeeResponse {
	return web.EmployeeResponse{
		Id:         employee.Id,
		Name:       employee.Name,
		Email:      employee.Email,
		Gender:     employee.Gender,
		Age:        employee.Age,
		Phone:      employee.Phone,
		TeamId:     employee.TeamId,
		RoleId:     employee.RoleId,
		IsVerified: employee.IsVerified,
	}
}

func ToEmployeeResponses(employees []domain.Employee) []web.EmployeeResponse {
	var employeeResponses []web.EmployeeResponse
	for _, employee := range employees {
		employeeResponses = append(employeeResponses, ToEmployeeResponse(employee))
	}
	return employeeResponses
}
