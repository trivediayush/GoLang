package main

import "fmt"

func main() {
	age := 15

	if age >= 18 {
		fmt.Println("You can Drive....")
	} else if age > 16 && age < 18 {
		fmt.Println("You can learn driving...")
	} else {
		fmt.Println("You are a kid...")
	}
}
