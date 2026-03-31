package main

import "fmt"

func main() {

	// age := 13

	// if age >= 18 {
	// 	fmt.Println("adult")
	// } else if age >= 12 {
	// 	fmt.Println("teen")
	// } else {
	// 	fmt.Println("child")
	// }

	// var role = "3"
	// var hasPermission = false

	// if role == "admin" && hasPermission == true {
	// 	fmt.Println("yes")
	// }

	if age := 15; age >= 18 {
		fmt.Println("adult", age)
	} else if age > 12 {
		fmt.Println("teen", age)

	}

	// [No ternary operator in go]
}
