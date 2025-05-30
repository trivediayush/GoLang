package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Print("its weekend")
	default:
		fmt.Printf("its workday")
	}
}