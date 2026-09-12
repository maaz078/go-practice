package main

import "fmt"

func worker(i int) {
	fmt.Printf("worker %d started\n", i)
	//some task happing there
	fmt.Printf("worker %d end\n", i)
}

func main() {
	for i := 1; i <= 3; i++ {
		go worker(i)
	}
	fmt.Println("worker task completed")
}
