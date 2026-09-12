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

	marks, exits := student["Math"]
	fmt.Println(marks)
	fmt.Println(exits)

	for student, mark := range student {
		fmt.Println(student, mark)
	}

	// students := make(map[string]int)

	// var name string
	// var mark int

	// fmt.Print("enter your name ")
	// fmt.Scan(&name)

	// fmt.Println("enter your mark ")
	// fmt.Scan(&mark)

	// students[name] = mark

	// fmt.Println(students)

	//Use a map to count how many times each number appears.
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4}

	frequency := make(map[int]int)

	for _, num := range numbers {
		frequency[num]++
	}

	for num, count := range frequency {
		fmt.Println(num, "→", count)
	}
}
