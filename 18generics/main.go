package main

import "fmt"

// func printSlice(items []int) {
// 	for _, i := range items {
// 		fmt.Println(i)
// 	}
// }

// func printSlice[T any](items []T) {
// 	for _, i := range items {
// 		fmt.Println(i)
// 	}
// }

func printSlice[T int | string](items []T) {
	for _, i := range items {
		fmt.Println(i)
	}
}

func main() {
	printSlice([]int{1, 2, 3, 4})
	stringss := []string{"golang", "js"}
	fmt.Println(stringss)
}
