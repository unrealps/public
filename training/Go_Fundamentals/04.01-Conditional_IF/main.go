package main

import (
	"fmt"
)

func main() {
	var age int
	println("Enter your age:")
	fmt.Scan(&age)

	if age < 18 {
		println("You are a minor.")
	} else if age >= 18 && age < 65 {
		println("You are an adult.")
	} else {
		println("You are a senior citizen.")
	}
}
