package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var n = make([]int, 2, 5)
	// var n := [] int {}
	fmt.Println(cap(n))
	n = append(n, 1)
	n = append(n, 2)
	n = append(n, 3)
	n = append(n, 4)
	fmt.Println(n)
	fmt.Println(cap(n)) // capacity doubles


	var n1 = []int {1,2,3,4}

	fmt.Println(n1[:3])

}
