package main

import "fmt"

func main() {
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("It's an integer.")
		case float64:
			fmt.Println("It's a float.")
		case string:
			fmt.Println("It's a string.")
		case bool:
			fmt.Println("It's a boolean.")
		default:
			fmt.Println("It's something else.")
		}
	}

	whoAmI("ayush")
}
