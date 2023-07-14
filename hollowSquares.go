package main

import "fmt"

func HollowSquares() {
	var sideLength int
	fmt.Print("Side Length : ")
	fmt.Scan(&sideLength)
	if sideLength >= 1 && sideLength <= 20 {
		lineNo := 0
		for lineNo < sideLength {
			if lineNo == 0 || lineNo == sideLength-1 {
				columnNo := 0
				for columnNo < sideLength {
					fmt.Print("*")
					columnNo++
				}
				fmt.Println()
			} else {
				columnNo := 0
				for columnNo < sideLength {
					if columnNo == 0 || columnNo == sideLength-1 {
						fmt.Print("*")
					} else {
						fmt.Print(" ")
					}
					columnNo++
				}
				fmt.Println()
			}
			lineNo++
		}
	} else {
		fmt.Println("input range Exceeded")
		return
	}

}
