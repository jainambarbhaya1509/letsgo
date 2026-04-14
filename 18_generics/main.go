package main

import "fmt"


// int | string | stack[T]
func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {
	// nums := []int{1, 2, 3}
	// myStack := stack[int]{
	// 	elements: nums,
	// }
	names := []string{"go", "dart", "py"}
	printSlice(names)
}
