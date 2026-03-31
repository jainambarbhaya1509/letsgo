package main

import "fmt"

const name string = "golang"

func main() {

	fmt.Println(name)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
