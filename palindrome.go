package main

import (
	"fmt"
)

func Palindrome() {
	fmt.Println("Please enter a five digit number")
	var number int

	fmt.Print("number : ")
	fmt.Scan(&number)
	if number/1000 > 10 && number/10000 < 10 {
		var first int
		var second int
		var third int
		var fourth int
		var fifth int
		var updatedNumber = number
		fifth = updatedNumber / 10000
		updatedNumber = (number - fifth*10000)
		fourth = updatedNumber / 1000
		updatedNumber = (updatedNumber - fourth*1000)
		third = updatedNumber / 100
		updatedNumber = (updatedNumber - third*100)
		second = updatedNumber / 10
		updatedNumber = (updatedNumber - second*10)
		first = updatedNumber / 1

		if fifth == first && second == fourth {
			fmt.Println("given number is a palindrome")
		} else {
			fmt.Println("given number is a not a palindrome")
		}

	} else {
		fmt.Println("given number doesnt match the required input")
	}
}
