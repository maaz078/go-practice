package main

import "fmt"

func main() {
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 1; i <= 20; i++ {
	// 	if i%2 != 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 != 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// for i := 1; i <= 5; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// var num int
	// fmt.Println("enter your number")
	// fmt.Scan(&num)

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(num, "x", i, "=", num*i)
	// }
	// var num int
	// fmt.Println("enter you table number: ")
	// fmt.Scan(&num)

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(num, "X", i, "=", num*i)
	// }

	// sum := 0
	// for i := 1; i <= 10; i++ {
	// 	sum = sum + i
	// }
	// fmt.Println("Sum = ", sum)

	// sum := 1
	// for i := 1; i <= 10; i++ {
	// 	sum = sum * i
	// }
	// fmt.Println(sum)

	//Write a Go program to print numbers from 1 to 100 using a for loop.

	// for i := 1; i<=100 ; i++{
	// 	fmt.Println(i)
	// }

	//Write a Go program to print numbers from 50 to 1 in reverse order.
	// for i := 50; i >= 1; i-- {
	// 	fmt.Println(i)
	// }

	//Write a Go program to print all even numbers between 1 and 50.
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//Write a Go program that takes a number from the user and prints its multiplication table from 1 to 10.
	// var n int
	// fmt.Println("enter your number")
	// fmt.Scan(&n)

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(n, "x", i, "=", n*i)
	// }

	//Write a Go program to calculate the sum of numbers from 1 to 100.
	// sum := 0
	// for i := 1; i <= 100; i++ {
	// 	sum = sum + i
	// }

	// fmt.Println(sum)

	//Write a Go program that takes a number from the user and calculates its factorial.
	var n int
	factorial := 1

	fmt.Println("enter your number")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		factorial = factorial * i
	}
	fmt.Println(factorial)
}
