package employees

import (
	"database/sql"
	"fmt"
	"shop/core"
)

type EmployeesStorage struct {
	DB *sql.DB
}

func (em EmployeesStorage) CreateEmployees(employees core.Employees) {
	query := "insert into employees(name) values ($1)"
	_, err := em.DB.Exec(query, employees.Name)
	if err != nil {
		fmt.Println("CreateEmployees, query error", err)
		return
	}
}

func (em EmployeesStorage) UpdateEmployees(employees core.Employees) {
	query := "update employees set name = $1 where id = $2; "
	_, err := em.DB.Exec(query, employees.Name, employees.ID)
	if err != nil {
		fmt.Println("UpdateEmployees, query erorr", err)
		return
	}
}

func (em EmployeesStorage) DeleteById(id int) {
	query := "delete from employees where id = $1"
	result, err := em.DB.Exec(query, id)
	if err != nil {
		fmt.Println("DeleteById, query erorr", err)

		return
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		return
	}
}

func (em EmployeesStorage) GetAllEmployees() []core.Employees {
	query := "select id, name from employees"
	row, err := em.DB.Query(query)
	if err != nil {
		fmt.Println("GetAllEmployees, query erorr", err)
		return []core.Employees{}
	}
	var employees []core.Employees

	for row.Next() {
		var employee core.Employees
		err := row.Scan(&employee.ID, &employee.Name)
		if err != nil {
			fmt.Println(" GetAllEmployees, Scan error", err)
			return []core.Employees{}
		}
		employees = append(employees, employee)

	}
	return employees
}
