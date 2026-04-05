package main

import "fmt"

func main() {
	var nums [4]bool
	fmt.Println(len(nums))
	nums[0] = true
	fmt.Println(nums)

	nums1 := [3]int{
		1, 2, 3,
	}

	fmt.Println(nums1)

	//2d array

	n := [8][2]int{
		{1, 2}, {3, 4},
		{1, 2}, {3, 4},
		{1, 2}, {3, 4},
		{1, 2}, {3, 4},
	}
	fmt.Println(n)
}
