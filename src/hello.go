package main

import (
	"fmt"
)

func main() {
	helloWorld := getHelloWorld()
	fmt.Println(helloWorld)
}

func getHelloWorld() string {
	return "Hello World!"
}
