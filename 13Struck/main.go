package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time
}

type Student struct {
	Name  string
	Age   int
	Marks float64
}

func main() {

	var firstOrder order = order{
		id:     "1",
		amount: 50.2,
		status: "recieved",
	}
	firstOrder.createdAt = time.Now()
	fmt.Println(firstOrder)

	// student1 := Student{
	// 	Name:  "Maaz",
	// 	Age:   19,
	// 	Marks: 12.3,
	// }

	// student2 := Student{
	// 	Name:  "Mokkaram",
	// 	Age:   21,
	// 	Marks: 12.3,
	// }

	// fmt.Println(student1.Name)
	// fmt.Println(student2)

	s := []Student{
		{
			Name:  "Saad",
			Age:   23,
			Marks: 23,
		},
		{
			Name:  "shareef",
			Age:   25,
			Marks: 89,
		},
		{
			Name:  "kazim",
			Age:   28,
			Marks: 34,
		},
	}
	fmt.Println(s)

	for _, s := range s {
		fmt.Println("Name:", s.Name)
		fmt.Println("Age:", s.Age)
		fmt.Println("Marks:", s.Marks)
	}

	s = append(s, Student{
		Name:  "Sahil",
		Age:   22,
		Marks: 88,
	})

	var student Student

	fmt.Print("Enter name: ")
	fmt.Scan(&student.Name)

	fmt.Print("Enter age: ")
	fmt.Scan(&student.Age)

	fmt.Print("Enter marks: ")
	fmt.Scan(&student.Marks)

	fmt.Println("\nStudent Details")
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("Marks:", student.Marks)

	language := struct {
		name   string
		isgood bool
	}{"maaz", true}
	fmt.Println(language)
}
