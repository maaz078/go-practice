package main

import "fmt"

func main() {

	// age := 20
	// if age > 18 {
	// 	fmt.Println("adult")
	// } else {
	// 	fmt.Println("not adult")
	// }

	// var mark int
	// fmt.Println("enter your marks...")
	// fmt.Scan(&mark)

	// if mark >= 90 {
	// 	fmt.Println("Grade A")
	// } else if mark >= 80 {
	// 	fmt.Println("Grade B")
	// } else if mark >= 60 {
	// 	fmt.Println("Grade C")
	// } else if mark >= 40 {
	// 	fmt.Println("Grade E")
	// } else {
	// 	fmt.Println("fail")
	// }

	// age := 20
	// if age >= 18 && age <= 60 {
	// 	fmt.Println("eligible")
	// }

	// day := "Sunday"
	// if day == "Sunday" || day == "Saturday" {
	// 	fmt.Println("weekend")
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
