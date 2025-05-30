package main

import "fmt"

func main() {
	var age = 15
	if age >= 18 {
		fmt.Println("You can vote...")
	}else{
		fmt.Println("You can not vote..")
	}
}