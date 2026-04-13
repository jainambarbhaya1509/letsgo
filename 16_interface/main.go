package main

import "fmt"

type paymenter interface{
	pay(amount float32)
}

type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}
type stripe struct{}


func (r razorpay) pay(amount float32) {
	fmt.Println("Razorpay: ", amount)
}
func (r stripe) pay(amount float32) {
	fmt.Println("Stripe: ", amount)
}

func main() {
	// stripePg := stripe{}
	razorpayPg := razorpay{}
	pg := payment{
		gateway: razorpayPg,
	}
	pg.makePayment(100)
}
