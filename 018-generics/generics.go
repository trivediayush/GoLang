package main

import "fmt"

// Generic function to print any type of slice
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

func main() {
    ints := []int{1, 2, 3}
    strings := []string{"hello", "world"}

    PrintSlice(ints)    // Prints: 1 2 3
    PrintSlice(strings) // Prints: hello world
}
