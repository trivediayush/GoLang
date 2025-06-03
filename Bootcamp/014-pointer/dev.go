package main

import "fmt"

func swapIPs(ip1, ip2 *string) {
	*ip1, *ip2 = *ip2, *ip1
}

func main() {
	ipA := "192.168.1.10"
	ipB := "10.0.0.5"

	fmt.Println("Before Swap:", ipA, ipB)

	swapIPs(&ipA, &ipB)

	fmt.Println("After Swap:", ipA, ipB)
}
