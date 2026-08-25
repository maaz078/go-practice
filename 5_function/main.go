package main

import "fmt"

// 1 func great() {
// 	fmt.Println("hello")
// }

//2 func great(name string) {
// 	fmt.Println("hello", name)
// }

//3 func great(a int, b int) {
// 	fmt.Println(a + b)
// }

//4 func great(a int, b int) int {
// 	return a + b
// }

//5 func great(a int, b int) int {
// 	return a * b
// }

//7 func add(a int, b int) int {
// 	return a + b
// }

//8 func calculate(a int, b int) (int, int) {
// 	sum := a + b
// 	diff := a - b

// 	return sum, diff
// }

//Write a function that accepts a name as a parameter and prints:
// func great(name string) {
// 	fmt.Println("hello", name)
// }

//Write a function add() that accepts two integers and returns their sum.
func great(a int, b int) int {
	return a + b
}
func main() {

	// great()

	// great("Maaz")

	// great(2, 3)

	// result := great(2, 3)
	// fmt.Println(result)

	// var a, b int
	// fmt.Println("enter your first number")
	// fmt.Scan(&a)

	// fmt.Println("enter your second number")
	// fmt.Scan(&b)

	// result := add(a, b)
	// fmt.Println("Sum= ", result)

	// sum, diff := calculate(23, 34)
	// fmt.Println("Sum: ", sum)
	// fmt.Println("Diff: ", diff)

	// great("maaz")
	result := great(2, 3)
	fmt.Println(result)

}
