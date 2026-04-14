package main

import "fmt"

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changed order status:", status)
}

func main() {
	changeOrderStatus(Received)
}
