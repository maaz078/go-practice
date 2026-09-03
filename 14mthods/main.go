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

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return float32(o.amount)
}

func main() {
	myOrder := order{
		id:        "1",
		amount:    100.1,
		status:    "pending",
		createdAt: time.Now(),
	}
	myOrder.changeStatus("lol")
	fmt.Println(myOrder.getAmount())

	language := struct {
		name   string
		isgood bool
	}{"maaz", true}
	fmt.Println(language)
}
