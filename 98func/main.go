package main

import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

func getlanguages() (string, string, string) {
	return "html", "css", "javascript"
}

func main() {
	// result := add(3, 4)
	// fmt.Println(result)

	fmt.Println(getlanguages())
}
