package main

import (
	"fmt"
	"strings"
)

func LongestSubstring() {
	fmt.Println("Please input a list of words you want to find out the common substring")

	words := []string{}
	for i := 1; i > 0; i++ {
		fmt.Println("To stop giving words type exit and press enter")

		var word string
		fmt.Printf("Input word number %v ", i)
		fmt.Scan(&word)
		if word == "exit" {
			break
		}
		words = append(words, word)
	}
	var minLength int
	index := 0
	for i, word := range words {
		if i == 0 {
			minLength = len(word)
		} else {
			if len(word) < minLength {
				minLength = len(word)
				index = i
			}
		}
	}

	var smallestWord = words[index]
	largestSubString := ""

	for k := 0; k < len(smallestWord); k++ {
		for i := k; i < len(smallestWord); i++ {

			var subString = smallestWord[k : i+1]
			subStringCheck := true
			for j := 0; j < len(words); j++ {
				if !strings.Contains(words[j], subString) {
					subStringCheck = false
				}
			}
			if subStringCheck {
				largestSubString = subString
			}
		}
		if largestSubString != "" {
			fmt.Println("the largest substring is : " + largestSubString)
			break
		}
	}

}
