package main

import "fmt"

func main() {
	student := map[string]int{
		"Math":    98,
		"English": 40,
		"Urdu":    39,
	}

	fmt.Println(student["Math"])

	//data add karna
	student["Marathi"] = 66
	fmt.Println(student)

	//data delete karna
	delete(student, "Urdu")
	fmt.Println(student)

	students := make(map[string]int)

	var name string
	var mark int

	fmt.Print("enter your name ")
	fmt.Scan(&name)

	fmt.Println("enter your mark ")
	fmt.Scan(&mark)

	students[name] = mark

	fmt.Println(students)
}
