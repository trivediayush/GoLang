package main

import "fmt"

// Define an enum for days of the week
type Day int

const (
    Sunday Day = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func main() {
    var today Day = Wednesday

    fmt.Println("Today is", today)       // Prints the integer value
    fmt.Println("Today is", today.String()) // Prints the string representation
}

// Adding a String method to print the enum as a string
func (d Day) String() string {
    days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

    if d < Sunday || d > Saturday {
        return "Unknown"
    }
    return days[d]
}
