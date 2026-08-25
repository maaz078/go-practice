package main

import (
	"fmt"
	"slices"
)

func main() {
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var num = make([]int, 2, 5)
	// num[0] = 1
	// num[1] = 2
	// num = append(num, 1)
	// num = append(num, 2)
	// fmt.Println(cap(num))
	// fmt.Println(num)

	//slice operater
	var nums = []int{1, 2, 3, 4}
	fmt.Println(nums[1:])
	fmt.Println(nums[1:3])

	//slice
	var num1 = []int{1, 2, 3}
	var num2 = []int{1, 2, 3}
	fmt.Println(slices.Equal(num1, num2))

	//2d slicer
	var num = [][]int{{1, 2}, {1, 2}}
	fmt.Println(num)

}
