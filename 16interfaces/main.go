package main

import "fmt"

type payment struct{}

func (p payment) makepayment(amount float32) {
	// getpaymentWT := rozarpay{}
	// getpaymentWT.pay(amount)

	stripeWT := stripe{}
	stripeWT.pay(amount)
}

type rozarpay struct{}

func (r rozarpay) pay(amount float32) {
	fmt.Println("making payment from rozarpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment from stripe", amount)
}

func main() {
	newpayment := payment{}
	newpayment.makepayment(200)
}
