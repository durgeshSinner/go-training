package main

import (
	"fmt"
	"sort"
)

func SortNumbers() {
	fmt.Println("Please input a list of numbers you want to sort")

	numbers := []int{}
	for i := 1; i > 0; i++ {
		fmt.Println("To stop giving words type exit and press enter")

		var number int
		fmt.Printf("Input word number %v ", i)
		_, err := fmt.Scan(&number)
		if err != nil {
			panic("input mismatch")
		}
		if number == 0 {
			break
		}
		numbers = append(numbers, number)
	}
	sort.Ints(numbers)

	fmt.Println("sorted order for the numbers is : ")
	for _, number := range numbers {
		fmt.Println(number)
	}
}
