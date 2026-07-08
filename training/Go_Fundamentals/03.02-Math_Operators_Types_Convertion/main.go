package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Convert string to int
	var str string
	var num int

	// String to int conversion
	println("Enter a number as a string:")
	fmt.Scan(&str)

	num, err := strconv.Atoi(str)
	if err != nil {
		println("Error converting string to int:", err)
		return
	}

	println("Converted number:", num)

	// Another away to convert string to int using ParseInt
	// Without explicit type conversion, ParseInt returns an int64
	num64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		println("Error converting string to int using ParseInt:", err)
		return
	}

	println("Converted number using ParseInt:", num64)

	// Another away to convert string to int using ParseInt
	// With explicit type conversion, ParseInt returns an int64
	f64 := float64(num64) // Convert int to float64
	println("Converted int to float64:", f64)
}
