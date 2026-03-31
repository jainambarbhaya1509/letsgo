package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("defaut")
	}

	// handles break automatically

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Workday")

	}

	//type switch

	whoAmI := func(i any) {
		switch i.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		default:
			fmt.Println("other")
		}
	}

	whoAmI("234")
	whoAmI(234)

}
