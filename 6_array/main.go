package main

import "fmt"

func main() {
	// numbers := [3]int{2, 3, 4}
	// fmt.Println(numbers)
	// fmt.Println(numbers[2])

	// numbers[1] = 9
	// fmt.Println(numbers)
	// fmt.Println(len(numbers))

	// num := [3]int{7, 8, 9}

	// for i := 0; i < len(num); i++ {
	// 	fmt.Println(num[i])
	// }

	var nums [5]int

	fmt.Println("Enter your fav number")
	fmt.Scan(&nums)

	for i := 0; i < 5; i++ {

		fmt.Println("enter your number")
		fmt.Scan(&nums[i])
	}
	fmt.Println("ArreyL: ", nums)

	number := [5]int{10, 20, 30, 40, 50}
	sum := 0
	for i := 0; i < len(number); i++ {
		sum += number[i]
	}
	fmt.Println("Sum: ", sum)
}
