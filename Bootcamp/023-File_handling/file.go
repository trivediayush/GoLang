package main

import (
	"fmt"
	"os"
)

func main() {
	// Write to a file
	data := []byte("Hello, file handling in Go!")
	err := os.WriteFile("example.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	// Read from the file
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:", string(content))

	// Clean up: delete the file
	err = os.Remove("example.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
	}
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current directory:", err)
    return
}   
    fmt.Println("Current working directory:", cwd)

}
