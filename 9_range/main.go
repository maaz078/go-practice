package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4, 5}
	// for i := 0; i < len(num); i++ {
	// 	fmt.Println(num[i])
	// }

	sum := 1
	for _, nums := range num {
		sum *= nums
		fmt.Println(sum)
	}

	// sum := 0
	// for i, v := range num {
	// 	sum = sum + v
	// 	fmt.Println(i, v)
	// }

	// fmt.Println("Sum: ", sum)

	// m := map[string]string{
	// 	"name": "Alice",
	// 	"age":  "30",
	// }
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	m := map[string]string{"name": "alice", "age": "30"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// for i, c := range "Maaz" {
	// 	fmt.Println(i, string(c))
	// }
}
