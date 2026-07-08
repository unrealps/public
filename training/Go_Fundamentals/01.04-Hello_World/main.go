package main

import "fmt"

// main is the program entry point.
// It delegates the actual output to a small helper for readability.
func main() {
	printGreeting()
}

// printGreeting prints the classic first message used in many examples.
func printGreeting() {
	const greeting = "Hello, World!"
	fmt.Println(greeting)
}
