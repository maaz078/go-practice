package main

import (
	"fmt"
	"strconv"
)

func main() {
	// currntTime := time.Now()
	// formated := currntTime.Format("2006/01/02, 3:04 PM Monday")
	// fmt.Println(formated)

	// layout_str := "2006-01-02"
	// dateStr := "2023-11-25"
	// formatedTime, _ := time.Parse(layout_str, dateStr)
	// fmt.Println(formatedTime)

	// newdate := currntTime.Add(24 * time.Hour)
	// fmt.Println(newdate)
	// formated_newdate := newdate.Format("2006/01/02 Monday")
	// fmt.Println(formated_newdate)

	// var a int = 2
	// var b float64 = float64(a)
	// fmt.Printf("data type of b %T", b)

	num := 31
	str := strconv.Itoa(num)
	fmt.Printf("%T", str)
}
