package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func getLang() (string, string, string) {
	return "go", "py", "dart"
}

func processIt(fn func(a int) int) {
	fn(1)
}

func processIt2() func(a int) int {
	return func(a int) int {
		return 2
	}
}
func main() {
	// res := add(3,5)
	l1, l2, _ := getLang()

	fmt.Println(l1, l2)

	fn := func(a int) int {
		return a
	}

	fn2 := processIt2()
	fmt.Println(fn2(2))
	processIt(fn)
}
