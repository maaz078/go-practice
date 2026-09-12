package main

import (
	"fmt"
)

// Write a function named greet that prints: Call the function from main().
func great() {
	fmt.Println("hellooo")
}

// Write a function that accepts a name as a parameter and prints:
func g(name string) {
	fmt.Println("helllooo", name)
}

// Write a function add() that accepts two integers and returns their sum.
func add(a int, b int) int {
	return a + b
}

// Calculator Functions Create separate functions for:add()subtract()multiply()divide()
func sub(a int, b int) int {
	return a - b
}

//Q7.⭐Multiple Return Values Write a function that accepts two integers and returns:

// Write a function named factorial() that accepts an integer and returns its factorial.
func factorial(n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
func man(a int, b int) (int, int, int) {
	sum := a + b
	diff := a - b
	product := a * b

	return sum, diff, product
}
func main() {
	great()
	g("Maaz")
	r := add(2, 3)
	fmt.Println(r)
	s := sub(2, 3)
	fmt.Println(s)

	sum, diff, product := man(5, 10)
	fmt.Println(sum)
	fmt.Println(diff)
	fmt.Println(product)

	fmt.Println(factorial(5))
}
