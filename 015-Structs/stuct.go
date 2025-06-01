package main

import (
	"fmt"
	"time"
)

// Define a struct type named 'order'
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

// Constructor-like function to create a new order
func newOrder(id string, amount float32, status string) *order {
	myorder := order{
		id:     id,       // use input parameters, not hardcoded
		amount: amount,
		status: status,
	}
	myorder.createdAt = time.Now() // set creation time
	return &myorder                // return a pointer to the new order
}

// Method with pointer receiver to change status
func (o *order) changestatus(status string) {
	o.status = status
}

// Method with value receiver to get order amount
func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// Define and initialize an anonymous struct (language info)
	language := struct {
		name   string
		isgood bool
	}{
		name:   "golang", // correct field initialization
		isgood: true,     // correct boolean value (was "ture" before)
	}

	fmt.Println("Language:", language)

	// Correct way to initialize using a constructor function
	myorder2 := newOrder("02", 30.57, "received")
	myorder2.changestatus("confirmed")
	fmt.Println("Order 2:", *myorder2) // dereference to print actual struct

	// Manual struct initialization
	myorder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}
	myorder.createdAt = time.Now()
	myorder.changestatus("confirmed")

	fmt.Println("Order 1:", myorder)
	fmt.Println("Order 1 Amount:", myorder.getAmount())
}
