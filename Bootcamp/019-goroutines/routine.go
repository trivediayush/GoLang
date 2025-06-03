package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    go sayHello() // starts a new goroutine
    time.Sleep(1 * time.Second) // wait to let goroutine finish
    fmt.Println("Main function finished")
}
