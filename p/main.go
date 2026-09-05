package main

import "fmt"

func main() {
	greet := func(a int, b int) int {
		return a + b
	}

	add := greet(90, 90)
	fmt.Println(add)

	var x, y int
	fmt.Println("enter x value: ")
	fmt.Scan(&x)
	fmt.Println("enter y value: ")
	fmt.Scan(&y)

	r := func(x int, y int) int {
		return x * y
	}
	s := r(x, y)
	fmt.Println(s)

	hell := func(a int, b int) (int, int) {
		sum := a + b
		diff := a - b
		return sum, diff
	}
	sum, diff := hell(12, 23)
	fmt.Println(sum, diff)
}
