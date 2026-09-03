package main

import "fmt"

type payment struct{}

func (p payment) makePayments(amount float64) {
	// rozarpayPaymentGW := rozarpay{}
	// rozarpayPaymentGW.pay(amount)

	stripePaymentGW := stripe{}
	stripePaymentGW.pay(amount)
}

type rozarpay struct{}

func (r rozarpay) pay(amount float64) {
	fmt.Println("making payment using rozarpay...", amount)
}

type stripe struct{}

func (s stripe) pay(amount float64) {
	fmt.Println("making payment using stripe...", amount)
}

func main() {

	newpayment := payment{}
	newpayment.makePayments(100)
}
