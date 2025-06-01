package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getlang() (string, string, bool) {
	return "golang", "python", true
}

func main() {
	result := add(4, 5)
	fmt.Println(result)
	lang1, lang2, lang3 := getlang()
	fmt.Println(lang1, lang2, lang3)
	lang4, lang5, _ := getlang()
	fmt.Println(lang4, lang5)
}