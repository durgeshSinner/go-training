package main

import "fmt"

type Employee struct {
	name   string
	hours  int
	salary int
}

var pricePerHour int = 200

func EmployeeSalary() {

	employeeCount := 0

	fmt.Println("Welcome to the Employee salary Calculator")
	for true {
		employeeCount++
		var employeeName string
		var hours int
		var salary int

		fmt.Printf("Please input employee %v details\n", employeeCount)
		fmt.Print("Name : ")
		fmt.Scan(&employeeName)
		fmt.Print("Hours : ")
		fmt.Scan(&hours)
		salary = hours * pricePerHour

		fmt.Printf("Hello %v you are earning %v\n", employeeName, salary)

	}
}
