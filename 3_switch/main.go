package main

import "fmt"

func main() {
	//Write a Go program that takes a number from the user and prints the corresponding day.

	// var day int
	// fmt.Println("enter your day number")
	// fmt.Scan(&day)

	// switch day {
	// case 1:
	// 	fmt.Println("monday")
	// case 2:
	// 	fmt.Println("tuesday")
	// case 3:
	// 	fmt.Println("wednessday")
	// case 4:
	// 	fmt.Println("thursday")
	// case 5:
	// 	fmt.Println("friday")
	// case 6:
	// 	fmt.Println("saturday")
	// case 7:
	// 	fmt.Println("sunday")
	// }

	// days := "Saturday"
	// switch days {
	// case "Saturday", "Sunday":
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Weekday")
	// }

	//Simple Calculator
	//Write a Go program that takes two numbers and an operator (+, -, *, /) from the user and calculates
	// the result using switch

	// var a, b int
	// var operater string

	// fmt.Println("Enter you first number...")
	// fmt.Scan(&a)

	// fmt.Println("Enter your operater like (+, -, *, /, %)")
	// fmt.Scan(&operater)

	// fmt.Println("Enter your second number...")
	// fmt.Scan(&b)

	// switch operater {
	// case "+":
	// 	fmt.Println("Result: ", a+b)
	// case "-":
	// 	fmt.Println("Result: ", a-b)
	// case "*":
	// 	fmt.Println("Result: ", a*b)
	// case "/":
	// 	fmt.Println("Result: ", a/b)
	// case "%":
	// 	fmt.Println("Result: ", a%b)
	// default:
	// 	fmt.Println("Result: ", a/b)
	// }

	//Write a Go program that takes a traffic light color as input.

	var color string
	fmt.Println("enter your traffic color...")
	fmt.Scan(&color)

	switch color {
	case "yellow":
		fmt.Println("go slow")
	case "green":
		fmt.Println("go fast")
	case "red":
		fmt.Println("stop!")
	}
}
