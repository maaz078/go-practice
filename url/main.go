package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("print url")
	myURL := "https://example.com:8080/path/to/resource?key1=value&key2"

	parseedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("can't parse url", err)
		return
	}
	fmt.Println(parseedURL)

	fmt.Println("Schema: ", parseedURL.Scheme)
	fmt.Println("Host: ", parseedURL.Host)
	fmt.Println("Path: ", parseedURL.Path)
	fmt.Println("RawQuery: ", parseedURL.RawQuery)

	parseedURL.Path = "/newPath"
	fmt.Println(parseedURL)

}
