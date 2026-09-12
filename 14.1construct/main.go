package main

import (
	"fmt"
	"time"
)

type order struct {
	id      string
	amount  float64
	status  string
	creatAt time.Time
}

func newOrde(id string, amount float64, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

func main() {

	myOrder := newOrde("1", 32.2, "true")

	fmt.Println(myOrder)
}
