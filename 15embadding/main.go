package main

import (
	"fmt"
	"time"
)

type info struct {
	fname   string
	lname   string
	surname string
}

type detail struct {
	email string
	phone string
}

type address struct {
	roomNo string
	area   string
}

type employee struct {
	id string
	info
	detail
	address
}

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
	employee := employee{
		id: "1",
	}
	employee.address = address{
		roomNo: "1",
		area:   "Mumbra",
	}
	employee.info = info{
		fname:   "Maaz",
		lname:   "Khan",
		surname: "MaazKhan",
	}
	employee.detail = detail{
		email: "Maaz@12",
		phone: "213244",
	}
	fmt.Println(employee)

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
