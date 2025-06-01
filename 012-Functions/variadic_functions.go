// A variadic function accepts a variable number of arguments of the same type. Internally, these are received as a slice.
// Useful when you don't know how many inputs will be passed (e.g., aggregating multiple metrics, summing integers, etc.).
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println("Sum:", sum(1, 2, 3))       // 6
    fmt.Println("Sum:", sum(10, 20, 30, 5)) // 65
}
