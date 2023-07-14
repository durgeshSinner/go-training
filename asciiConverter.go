package main

import (
	"encoding/hex"
	"fmt"
)

func asciiConverter() {
	inputString := "IronMan part 3"
	hx := hex.EncodeToString([]byte(inputString))
	fmt.Println("String to Hex Golang example")
	fmt.Println()
	fmt.Println(inputString + " ==> " + hx)

}
