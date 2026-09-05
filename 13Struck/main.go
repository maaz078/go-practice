package main

import (
	"fmt"
	"time"
)

type student struct {
	Name string
	age  int
	mark float32
}

//	type order struct {
//		id        string
//		amount    float64
//		status    string
//		createdAt time.Time
//	}
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o *order) getamount() float32 {
	return o.amount
}

func main() {

	myorder := order{
		id:     "1",
		amount: 123.1,
		status: "pending",
	}
	myorder.changeStatus("confirmed")

	fmt.Println(myorder.getamount())
	fmt.Println(myorder)

	// var student1 student
	// student1.Name = "Maaz"
	// student1.age = 19
	// student1.mark = 99.8
	// fmt.Println(student1)
	// // fmt.Println(student1.Name)

	// myorder2.createdAt = time.Now()
	// fmt.Println(myorder2)
}
