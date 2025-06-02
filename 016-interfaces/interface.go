package main

import "fmt"

// paymentGateway interface
type paymentGateway interface {
	pay(amount float32)
}

// payment struct
type payment struct {
	gateway paymentGateway
}

// makePayment uses the provided paymentGateway to make a payment
func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

// razorpay struct implements the paymentGateway interface
type razorpay struct{}

// pay method for razorpay
func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using Razorpay...", amount)
}

// main function
func main() {
	// Create a Razorpay gateway
	rp := razorpay{}

	// Inject the Razorpay gateway into payment
	p := payment{gateway: rp}

	// Make a payment
	p.makePayment(100)
}
