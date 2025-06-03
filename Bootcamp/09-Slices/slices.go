// slices are dynamic array, when you dont know array size then we use it

package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(nums)         // []
	fmt.Println(nums == nil)  // true
	fmt.Println(len(nums))    // 0

	nums = append(nums, 1)
	fmt.Println(nums)         // [1]

	var num2 = make([]int, 4, 10) // valid: len <= cap
	fmt.Println(num2)
 // avoid unused variable warning

//  slice operator 

// 2d slices 

	var num2d = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(num2d)

}
