package main

import "fmt"

func main() {
	i := 1
	for i <= 6 {
		if i == 3 {
			i++
			continue
		}
		fmt.Println(i)
		i++
	}
}
