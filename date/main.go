package main

import (
	"fmt"
)

func a(z *int) {
	y := z
	*y = 50
}

func main() {
	// currentdate := time.Now()
	// fmt.Println(currentdate)

	// formatted := currentdate.Format("02-01-2006")
	// fmt.Println(formatted)

	aa := 20
	a(&aa)

	fmt.Println(aa)
}
