package main

import "fmt"

func printSlice[T int | string | bool](item []T) {
	for _, v := range item {
		fmt.Println(v)
	}
}

// func printStringSlice(item []string) {
// 	for _, v := range item {
// 		fmt.Println(v)
// 	}
// }

type stack[T any] struct {
	elements []T
}

func main() {

	// num := []int{1, 2, 3}

	num1 := []string{"maaz", "khan"}
	printSlice(num1)
	// printStringSlice(num1)

	myStack := stack[string]{
		elements: []string{"maaz", "khan"},
	}
	fmt.Println(myStack)

}
