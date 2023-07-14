package main

import (
	"fmt"
	"sort"
	"strconv"
)

func SortNumbers() {
	fmt.Println("Please input a list of numbers you want to sort")

	numbers := []int{}
	for i := 1; i > 0; {
		var number string
		fmt.Println("To stop giving numbers type X")
		fmt.Printf("Input word number %v ", i)
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("Input Mismatch")
			continue
		}
		if number == "X" {
			fmt.Println("Exiting the application")
			break
		}
		numberAsInt, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("invalid input try again")
			continue
		}
		numbers = append(numbers, int(numberAsInt))
		i++
	}
	sort.Ints(numbers)

	fmt.Println("sorted order for the numbers is : ")
	for _, number := range numbers {
		fmt.Println(number)
	}

}
