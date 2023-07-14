package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello welcome to go training course Assesment")
	fmt.Println("Which Question you want to attempt running ? ")
	var questionNumber int
	fmt.Scan(&questionNumber)
	switch questionNumber {
	case 1:
		EmployeeSalary()
	case 2:
		HollowSquares()
	case 3:
		Palindrome()
	case 4:
		asciiConverter()
	case 5:
		LongestSubstring()
	case 6:
		SortNumbers()

	default:
		fmt.Println("Sorry you havent selected anything")
	}
}
