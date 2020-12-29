package main

import "fmt"

func main() {
	go helloWorld()
}

func helloWorld() {
	fmt.Println("Hello world")
}
