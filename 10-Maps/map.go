package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)
	m["name"] = "golang"
	m["area"] = "backend"

	fmt.Println(m["name"],m["area"])

	n := make(map[string]int)
	n["age"] = 30

	fmt.Println(n["age"], n["phone"]) //phoen return zero becasue phone key does not exists in 'n' map and its type is int so it return 0

	mymap := map[string]int{"price":40000, "phone":24}

	fmt.Println(mymap)

	v, ok := mymap["phone"]
	fmt.Println(v)
	if ok {
		fmt.Println("All Ok")
	} else{
		fmt.Println("Not Ok")
	}

	m1 := map[string]int{"price":40000, "name": 7}
	m2 := map[string]int{"price":40000, "name": 7}

	fmt.Println(maps.Equal(m1, m2))
}

