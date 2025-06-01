package main

import "fmt"

func main() {
	// normal way 
	nums := []int{6, 7, 8}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// using range
	for j, num := range nums{
		fmt.Println(num, j)
	}

	// unicode, rune 
}