package main

import (
	"fmt"
)

func main() {
	var name string
	var age int

	// Asking for user input
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	// Output the input back to the user
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
