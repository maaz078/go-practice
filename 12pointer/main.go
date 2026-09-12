package main

import "fmt"

func changeNum(num *int) {
	*num = 90
}

// func changeAge(age *int) {
// 	*age = 25
// }

func main() {

	age := 18
	old := &age
	fmt.Println("value", age)

	fmt.Println("address", old)
	fmt.Println("value through pointer: ", *old)

	a := 10
	b := &a
	*b = 30
	fmt.Println(a)

	num := 1
	changeNum(&num)
	fmt.Println("change num after pointer...", num)

	// age := 20

	// changeAge(&age)

	// fmt.Println(age)

}
