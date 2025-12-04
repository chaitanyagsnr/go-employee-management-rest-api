package service

import (
	"database/sql"
	"employee-management-service/dto"

	_ "github.com/lib/pq"
)

func getDBConnection() *sql.DB {
	connString := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	return db
}

func ListEmployees() []dto.Employee {
	db := getDBConnection()
	defer db.Close()
	rows, err := db.Query("Select id, name, salary from Employee")
	if err != nil {
		panic(err)
	}

	var employees []dto.Employee
	for rows.Next() {
		var employee dto.Employee
		rows.Scan(&employee.Id, &employee.Name, &employee.Salary)
		employees = append(employees, employee)
	}

	return employees
}

func AddEmployee(employee dto.Employee) {
	db := getDBConnection()
	defer db.Close()
	db.Exec("Insert into Employee (id, name, salary) values ($1, $2, $3)", employee.Id, employee.Name, employee.Salary)
}
