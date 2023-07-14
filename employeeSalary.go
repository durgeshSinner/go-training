package main

import (
	"fmt"
)

type Employee struct {
	name   string
	hours  int
	salary int
}

var pricePerHour int = 200

func EmployeeSalary() {

	var employees []Employee
	fmt.Println("Welcome to the Employee salary Calculator")
	for i := 0; i < 3; {

		var employeeName string
		var hours int
		var salary int

		fmt.Printf("Please input employee %v details\n", i+1)
		fmt.Print("Name : ")
		fmt.Scan(&employeeName)
		fmt.Print("Hours : ")
		fmt.Scan(&hours)
		if hours > 40 {
			salary = 40*pricePerHour + ((hours - 40) * pricePerHour / 2)
		} else {
			salary = hours * pricePerHour
		}

		var employee = Employee{
			name:   employeeName,
			hours:  hours,
			salary: salary,
		}
		employees = append(employees, employee)
		fmt.Printf("Hello %v you are earning %v\n", employee.name, employee.salary)
		i++

	}
	fmt.Println("All employees details : ")
	for _, employee := range employees {
		fmt.Printf("Employee Name %v worked for %v and earned %v \n", employee.name, employee.hours, employee.salary)

	}
}
