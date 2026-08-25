package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5}
	number = append(number, 7)
	fmt.Println(number)
	fmt.Println(len(number))
	fmt.Println(cap(number))

	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}
}
