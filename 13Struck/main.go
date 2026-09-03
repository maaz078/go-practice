package main

import (
	"fmt"
	"time"
)

// type student struct {
// 	Name string
// 	age  int
// 	mark float64
// }

type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time
}

func main() {
	// var student1 student

	// student1.Name = "Maaz"
	// student1.age = 29
	// student1.mark = 99.8

	// fmt.Println(student1)
	// fmt.Println(student1.Name)

	myorder := order{
		id:        "1",
		amount:    100.2,
		status:    "pending",
		createdAt: time.Now(),
	}
	fmt.Println(myorder)
	fmt.Println(myorder.status)

	myorder2 := order{
		id:     "2",
		amount: 232.3,
		status: "Complted",
	}

	myorder2.createdAt = time.Now()
	fmt.Println(myorder2)
}
