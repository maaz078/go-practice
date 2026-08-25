package main

import (
	"fmt"
	"maps"
)

func main() {
	// m := make(map[string]string)
	// m["name"] = "Maaz"
	// m["surname"] = "Khan"
	// fmt.Println(m["name"], m["surname"])
	// fmt.Println(len(m))

	// delete(m, "surname")
	// fmt.Println(m)

	// m := map[string]int{"Price": 30, "phone": 40}
	// v, ok := m["Price"]
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("all good")
	// } else {
	// 	fmt.Println("bad")
	// }

	m1 := map[string]int{"Price": 30, "phone": 40}
	m2 := map[string]int{"Price": 30, "phone": 40}

	fmt.Println(maps.Equal(m1, m2))
}
