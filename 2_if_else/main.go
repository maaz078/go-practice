package main

import "fmt"

func main() {
	// age := 20
	// if age > 18 {
	// 	fmt.Println("you are adult")
	// } else {
	// 	fmt.Println("you are not adult")
	// }

	// var marks int
	// fmt.Println("Enter your marks...")
	// fmt.Scan(&marks)

	// if marks >= 90 {
	// 	fmt.Println("Great A")
	// } else if marks >= 80 {
	// 	fmt.Println("Great B")
	// } else if marks >= 70 {
	// 	fmt.Println("great C")
	// } else if marks >= 60 {
	// 	fmt.Println("great D")
	// } else if marks >= 40 {
	// 	fmt.Println("Great F")
	// } else {
	// 	fmt.Println("Fail")
	// }

	// age := 20
	// if age >= 18 && age <= 60 {
	// 	fmt.Println("Eligible")
	// }

	// day := "Sunday"
	// if day == "saturday" || day == "Sunday" {
	// 	fmt.Println("Weekend")
	// }

	// isStudent := false

	// if !isStudent {
	// 	fmt.Println("Not student")
	// }

	//1
	// var num int
	// fmt.Println("enter your number...")
	// fmt.Scan(&num)

	// if num > 0 {
	// 	fmt.Println("Positive")
	// } else if num < 0 {
	// 	fmt.Println("Negative")
	// } else {
	// 	fmt.Println("zero")
	// }

	//2
	var age int
	fmt.Println("enter your age...")
	fmt.Scan(&age)

	if age <= 0 || age <= 12 {
		fmt.Println("child")
	} else if age <= 13 || age <= 19 {
		fmt.Println("teenage")
	} else if age <= 20 || age <= 59 {
		fmt.Println("adult")
	} else {
		fmt.Println("too old")
	}
}
