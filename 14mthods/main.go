package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time
}

// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() (float64, string) {
	return o.amount, o.id
}

func main() {
	myOrder := order{
		id:        "1",
		amount:    100.1,
		status:    "pending",
		createdAt: time.Now(),
	}

	// myOrder.changeStatus("lol")
	myOrder.changeStatus("canceles")
	fmt.Println(myOrder.getAmount())
	fmt.Println(myOrder)

}
