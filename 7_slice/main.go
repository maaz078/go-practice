package main

import "fmt"

func main() {
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var num = make([]int, 2, 5)
	// num[0] = 1
	// num[1] = 2
	// num = append(num, 1)
	// num = append(num, 2)
	// fmt.Println(cap(num))
	// fmt.Println(num)

	//slice operater
	// var nums = []int{1, 2, 3, 4}
	// fmt.Println(nums[1:])
	// fmt.Println(nums[1:3])

	// //slice
	// var num1 = []int{1, 2, 3}
	// var num2 = []int{1, 2, 3}
	// fmt.Println(slices.Equal(num1, num2))

	// //2d slicer
	// var num = [][]int{{1, 2}, {1, 2}}
	// fmt.Println(num)

	// number := []int{1, 2, 3, 4, 5}
	// number = append(number, 7)
	// fmt.Println(number)
	// fmt.Println(len(number))
	// fmt.Println(cap(number))

	// for i := 0; i < len(number); i++ {
	// 	fmt.Println(number[i])
	// }
	//Write a Go program that prints every element of a slice using the range keyword.
	// number := []int{1, 2, 3, 4, 5}
	// for _, v := range number {
	// 	fmt.Println(v)
	// }

	//Write a Go program to calculate the sum of all elements in:
	// number := []int{1, 2, 3, 4, 5}
	// sum := 0
	// for i := 0; i < len(number); i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// Write a Go program to find the largest number in a slice.
	// a := []int{1, 2, 3, 4, 5}
	// max := a[0]
	// for i := 1; i < len(a); i++ {
	// 	if a[i] > max {
	// 		max = a[i]
	// 	}

	// }
	// fmt.Println(max)

	//Write a Go program that asks the user how many numbers they want to enter,
	//  creates a slice of that size, takes the numbers as input, and prints the slice.

	// var n int
	// println("how many number you want to enter?")
	// fmt.Scan(&n)

	// num := make([]int, n)

	// for i := 0; i < n; i++ {
	// 	fmt.Println("enter number: ")
	// 	fmt.Scan(&num[i])
	// }
	// fmt.Println("Slice: ", num)

	//Write a Go program that searches for a given number in a slice.
	num := []int{10, 20, 30, 40, 50}
	num = append(num[:2], num[3:]...)
	fmt.Println(num)
}
