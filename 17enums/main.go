package main

import "fmt"

// type OrderStatus int

// const (
// 	received OrderStatus = iota
// 	confirmed
// 	PrePared
// 	Delivered
// )

type OrderStatus string

const (
	received  OrderStatus = "received"
	confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to..", status)
}

func main() {
	changeOrderStatus(Delivered)

}
