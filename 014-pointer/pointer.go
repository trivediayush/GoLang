package main

import "fmt"

func updateValue(x *int) {
	*x = *x + 10 // dereference and modify the value at the address
}

func main() {
	num := 5
	fmt.Println("Before:", num)

	updateValue(&num) // pass the address of num

	fmt.Println("After:", num)
}
