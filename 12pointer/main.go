package main

import "fmt"

// func changenum(num *int) {
// 	*num = 3
// 	fmt.Println("In changeNum", *num)
// }

func changeAge(age *int) {
	*age = 25
}

func main() {

	age := 20

	changeAge(&age)

	fmt.Println(age)

	// age := 18
	// old := &age
	// fmt.Println("value", age)

	// fmt.Println("address", old)
	// fmt.Println("value through pointer: ", *old)

	// a := 10
	// b := &a
	// *b = 20
	// fmt.Println(a)

	// num := 1
	// changenum(&num)
	// fmt.Println("after change num in mai,", num)

}
