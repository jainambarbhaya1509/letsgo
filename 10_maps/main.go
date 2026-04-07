package main

import "fmt"

func main(){
	
	m := make(map[string]string)

	m["1"] = "1"
	m["2"] = "2"

	fmt.Println(m)

	fmt.Println(m["1"])
	delete(m, "2")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	// if key does not exist in the map then it return zero value

}