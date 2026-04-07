package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, i := range nums {
		total += i
	}
	return total
}
func main() {
	nums := []int{3, 4, 5, 6}
	res := sum(nums...)
	fmt.Println(res)
}
