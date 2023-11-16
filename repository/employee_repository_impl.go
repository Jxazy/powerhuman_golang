package repository

import (
	"context"
	"database/sql"
	"errors"
	"jxazy/powerhuman_golang/helper"
	"jxazy/powerhuman_golang/model/domain"
)

type EmployeeRepositoryImpl struct {
}

func NewEmployeeRepositoryImpl() EmployeeRepository {
	return &EmployeeRepositoryImpl{}
}

func (EmployeeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee {
	SQL := "insert into employees(name, email, gender, age, phone, team_id, role_id, is_verified) values (?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, employee.Name, employee.Email, employee.Gender, employee.Age, employee.Phone, employee.TeamId, employee.RoleId, employee.IsVerified)
	helper.CheckError(err)

	id, err := result.LastInsertId()
	helper.CheckError(err)

	employee.Id = int(id)

	return employee
}

func (EmployeeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee {
	SQL := "update employees SET name = ?, email = ?, gender = ?, age = ?, phone = ?, team_id = ?, role_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, employee.Name, employee.Email, employee.Gender, employee.Age, employee.Phone, employee.TeamId, employee.RoleId, employee.Id)
	helper.CheckError(err)
	return employee
}

func (EmployeeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, employee domain.Employee) {
	SQL := "delete from employees where id = ?"

	_, err := tx.ExecContext(ctx, SQL, employee.Id)
	helper.CheckError(err)
}

func (EmployeeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, employeeId int) (domain.Employee, error) {
	SQL := "select id,name,email, gender, age, phone, team_id, role_id, is_verified from employees where id = ?"

	rows, err := tx.QueryContext(ctx, SQL, employeeId)
	helper.CheckError(err)
	defer rows.Close()

	employee := domain.Employee{}
	if rows.Next() {
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Gender, &employee.Age, &employee.Phone, &employee.TeamId, &employee.RoleId, &employee.IsVerified)
		helper.CheckError(err)
		return employee, nil
	} else {
		return employee, errors.New("employee is not found")
	}
}

func (EmployeeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Employee {
	SQL := "select id,name,email, gender, age, phone, team_id, role_id, is_verified from employees"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.CheckError(err)
	defer rows.Close()

	var employees []domain.Employee
	for rows.Next() {
		employee := domain.Employee{}
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Gender, &employee.Age, &employee.Phone, &employee.TeamId, &employee.RoleId, &employee.IsVerified)
		helper.CheckError(err)
		employees = append(employees, employee)
	}

	return employees
}
