// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sayHi() {
// 	fmt.Println("Hiiiiii")
// 	time.Sleep(2000 * time.Millisecond)
// 	fmt.Println("after 1 second")
// }

// func sayHello() {
// 	fmt.Println("HEllloooo")
// }

// func main() {
// 	fmt.Println("learning concencey....")
// 	go sayHi()
// 	go sayHello()

// 	// time.Sleep(3000 * time.Millisecond)
// 	fmt.Println("hhhhhh")
// }

package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task doint", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
}
