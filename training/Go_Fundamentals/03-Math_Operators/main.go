package main

import (
	"fmt"
)

func main() {
	var a int
	var b int

	fmt.Println("Math Operators in Go")
	fmt.Println("--------------------")
	fmt.Print("Insert the first number: ")
	fmt.Scan(&a)
	fmt.Print("Insert the second number: ")
	fmt.Scan(&b)

	// Addition
	fmt.Println("Addition")
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	// Subtraction
	fmt.Println("Subtraction")
	diff := a - b
	fmt.Printf("%d - %d = %d\n", a, b, diff)

	// Multiplication
	fmt.Println("Multiplication")
	product := a * b
	fmt.Printf("%d * %d = %d\n", a, b, product)

	// Division
	fmt.Println("Division")
	if b != 0 {
		quotient := a / b
		fmt.Printf("%d / %d = %d\n", a, b, quotient)
	} else {
		fmt.Println("Cannot divide by zero")
	}

	// Modulus
	fmt.Println("Modulus")
	if b != 0 {
		remainder := a % b
		fmt.Printf("%d %% %d = %d\n", a, b, remainder)
	} else {
		fmt.Println("Cannot perform modulus by zero")
	}
}
