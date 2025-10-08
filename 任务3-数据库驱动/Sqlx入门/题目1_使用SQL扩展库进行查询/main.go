package main

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id         int
	Name       string
	Department string
	Salary     float64
}

/*
使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
并将结果映射到一个自定义的 Employee 结构体切片中。
*/
func getEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	stmt, err := db.Preparex("SELECT id, name, department,salary FROM employees WHERE department = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Queryx(department)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employeeSlice := make([]Employee, 0)
	for rows.Next() {
		var e Employee
		err := rows.StructScan(&e)
		if err != nil {
			log.Fatal(err)
		}
		employeeSlice = append(employeeSlice, e)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return employeeSlice, nil
}

/*
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
func getMaxSalaryEmployee(db *sqlx.DB) (*Employee, error) {
	e := Employee{}
	err := db.Get(&e,
		"SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func main() {
	db := sqlx.MustConnect("mysql", "root:123456@tcp(localhost:3306)/go_study")
	defer db.Close()

	employees, err := getEmployeesByDepartment(db, "技术部")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range employees {
		log.Printf("%+v\n", e)
	}

	e, err := getMaxSalaryEmployee(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Max Salary Employee: %+v\n", *e)
}
