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
	amount    float64
	status    string
	createdAt time.Time
	customer
}

func main() {

	// newcustomer := customer{
	// 	name:  "Maaz",
	// 	phone: "1234567890",
	// }

	newOrder := order{
		id:     "1",
		amount: 100.2,
		status: "pending",

		customer: customer{
			name:  "Maaz",
			phone: "1234567890",
		},
	}
	newOrder.createdAt = time.Now()
	newOrder.customer.name = "khan"
	fmt.Println(newOrder)
	fmt.Println(newOrder.customer)
}
