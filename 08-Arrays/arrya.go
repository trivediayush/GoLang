package main

import "fmt"

func main(){
	var nums = [4]int{1,2,3,4}
	nums2 := [...]int{1, 2, 3, 4}

	fmt.Println(len(nums))
	fmt.Println(nums[2])
	fmt.Println(nums2)
	fmt.Print(nums)

	var mybool [4]bool 
	mybool[0] = true
	fmt.Println(mybool)
	
	var mystring [4]string
	mystring[0] = "hello"
	fmt.Println(mystring)
	
	// 2D Array
	num2d := [2][2]int{{3,4},{5,6}}
	fmt.Println(num2d)
}