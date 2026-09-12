package main

import "fmt"

func main() {
	//Write a Go program to create an integer array containing:
	numbers := [3]int{2, 3, 4}
	fmt.Println(numbers)
	fmt.Println(numbers[2:])

	//Change the third element to 100 and print the array.
	numbers[1] = 30
	fmt.Println(numbers)

	//Create an array of 6 integers and print its length using len().
	// numbers[1] = 9
	// fmt.Println(numbers)
	// fmt.Println(len(numbers))

	//Write a Go program that prints every element of an array using a for loop.

	num := [3]int{7, 8, 9}
	for i := 0; i < (len(num)); i++ {
		fmt.Println(num[i])
	}

	//Write a Go program to calculate the sum of all elements in:
	sum := 0
	nums := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	fmt.Println(sum)

	//Write a Go program to find the largest number in an integer array.
	lar := [5]int{10, 23, 1, 3, 43}
	max := lar[0]
	for i := 0; i < len(lar); i++ {
		if lar[i] > max {
			max = lar[i]
		}
		fmt.Println(max)
	}

	//Write a Go program that takes 5 integers from the user, stores them in an array, and then prints the array.

	var numss [5]int

	fmt.Println("Enter your fav number")
	fmt.Scan(&numss)

	for i := 0; i < 5; i++ {

		fmt.Println("enter your number")
		fmt.Scan(&numss[i])
	}
	fmt.Println(numss)

	// number := [5]int{10, 20, 30, 40, 50}
	// sum := 0
	// for i := 0; i < len(number); i++ {
	// 	sum += number[i]
	// }
	// fmt.Println("Sum: ", sum)

	//Write a Go program that takes an array and a number from the user and checks whether that number exists in the array.
}
