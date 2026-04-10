package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

func (o *order) changeStatus(newStatus string) {
	o.status = newStatus
}

func main() {

	myOrder := order{
		id:     "1",
		amount: 123.321,
		status: "pending",
	}
	myOrder.createdAt = time.Now()

	myOrder.changeStatus("received")

	fmt.Println(myOrder)

	order1 := newOrder("1", 123, "paid")
	fmt.Println(order1)

	lang := struct {
		name   string
		isGood bool
	}{
		name:   "golang",
		isGood: true,
	}
	fmt.Println(lang)

	newCust := customer{
		name:  "abc",
		phone: "3211233211",
	}
	newOrder := order{
		id:       "1",
		amount:   123.321,
		status:   "pending",
		customer: newCust,
	}
	fmt.Println(newOrder)

}
