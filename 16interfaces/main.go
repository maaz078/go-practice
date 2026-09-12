// package main

// import "fmt"

// type payment struct{}

// func (p payment) pay(amount float32) {
// 	makePaymentGW := rozarpay{}
// 	makePaymentGW.pay(amount)
// }

// type rozarpay struct{}

// func (r rozarpay) pay(amount float32) {
// 	fmt.Println("making payment using rozarpay..", amount)
// }

// type stript struct{}

// func (s stript) pay(amount float32) {
// 	fmt.Println("making payment using stript...", amount)
// }
// func main() {
// 	newpayment := payment{}
// 	newpayment.pay(2000)
// }

// package main

// import "fmt"

// type payment struct {
// 	gateway razorpay
// }

// func (p payment) pay(amount float32) {
// 	p.gateway.pay(amount)
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) {
// 	fmt.Println("making payment using razarpay...", amount)
// }

// func main() {
// 	rozarpaymentGW := razorpay{}
// 	newpayment := payment{
// 		gateway: rozarpaymentGW,
// 	}
// 	newpayment.pay(34433)
// }

package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) pay(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razarpay...", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}
func main() {
	rozarpaymentGW := stripe{}
	newpayment := payment{
		gateway: rozarpaymentGW,
	}
	newpayment.pay(34433)
}

// package main

// import "fmt"

// type Animal interface {
//     Sound()
// }

// type Dog struct{}

// func (d Dog) Sound() {
//     fmt.Println("Dog says Woof")
// }

// type Cat struct{}

// func (c Cat) Sound() {
//     fmt.Println("Cat says Meow")
// }

// func main() {

//     var animal Animal

//     animal = Dog{}
//     animal.Sound()

//     animal = Cat{}
//     animal.Sound()
// }
