package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // mark this goroutine done when function exits
    fmt.Printf("Worker %d started\n", id)
    // do some work here
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1) // increment WaitGroup counter
        go worker(i, &wg)
    }
    wg.Wait() // wait for all goroutines to finish
    fmt.Println("All workers done")
}
