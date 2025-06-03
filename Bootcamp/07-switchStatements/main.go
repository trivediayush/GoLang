package main

import "fmt"

func main() {
	i := 5
	switch i {
	case 1:
		fmt.Print("one")
	case 2:
		fmt.Print("Two")
	case 3:
		fmt.Print("three")
	default:
		fmt.Print("others")
	}
}