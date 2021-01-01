package main

import (
	"fmt"
)

func main() {

	// c := make(chan string)
	// writing into the channel is not blocked until the capacity is reached (3)
	c := make(chan string, 3) // buffer channel that can hold 3 elements

	fmt.Println("Sending to channel")
	// side go routine:
	go func(input chan string) {
		fmt.Println("sending 'hello 1' to channel")
		input <- "hello 1"

		fmt.Println("sending 'hello 2' to channel")
		input <- "hello 2"

		fmt.Println("sending 'hello 3' to channel")
		input <- "hello 3"

		fmt.Println("sending 'hello 4' to channel")
		input <- "hello 4" // deadlock because capacity of buffer channel is full
	}(c)

	// main go routine: execution continues
	fmt.Println("Sending from channel")
	for greeting := range c {
		fmt.Println("\nGreeting received")
		fmt.Println(greeting)
	}
}
