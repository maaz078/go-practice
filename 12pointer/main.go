package main

import "fmt"

// func changenum(num *int) {
// 	*num = 3
// 	fmt.Println("In changeNum", *num)
// }

func main() {

	age := 18
	old := &age
	fmt.Println("value", age)

	fmt.Println("address", old)
	fmt.Println("value through pointer: ", *old)

	a := 10
	b := &a
	*b = 20
	fmt.Println(a)

	// num := 1
	// changenum(&num)
	// fmt.Println("after change num in mai,", num)

}
