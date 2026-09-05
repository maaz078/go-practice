package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.Create("example.text")
	// if err != nil {
	// 	fmt.Println("this is error while crating the file")
	// 	return
	// }
	// defer file.Close()

	// content := "hellloooo"
	// bytess, errors := io.WriteString(file, content)
	// fmt.Println("bytes", bytess)
	// if errors != nil {
	// 	fmt.Println("this is error while adding content inside the file", errors)
	// 	return

	// }
	// file, err := os.Open("example.text")
	// if err != nil {
	// 	fmt.Println("error while reading the file")
	// 	return
	// }

	// buffer := make([]byte, 1024)
	// for {
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("error while reading the file")
	// 	}
	// 	fmt.Println(string(buffer[:n]))
	// }

	// fmt.Println("all good")

	content, err := os.ReadFile("example.text")
	if err != nil {
		fmt.Println("error reading file")
		return
	}
	fmt.Println(string(content))
}
