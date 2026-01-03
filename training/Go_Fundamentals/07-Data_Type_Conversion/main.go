package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input1 string
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input1)

	// Convert string to int
	intValue, err := strconv.Atoi(input1)
	if err != nil {
		fmt.Println("Error converting to int:", err)
		return
	}

	// Convert string to float64
	floatValue, err := strconv.ParseFloat(input1, 64)
	if err != nil {
		fmt.Println("Error converting to float64:", err)
		return
	}

	// Display results
	fmt.Printf("You entered: %s\n", input1)
	fmt.Printf("Converted to int: %d\n", intValue)
	fmt.Printf("Converted to float64: %.2f\n", floatValue)

}
