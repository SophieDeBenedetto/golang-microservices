package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go helloWorld(c)
	fmt.Println("WHEN")
	greeting := <-c
	fmt.Println(greeting)
	// go helloWorld(greeting)
	// fmt.Println("Bye")
	// time.Sleep(1 * time.Millisecond)
}

func helloWorld(c chan string) {
	fmt.Println("IN HELLO")
	c <- "hello"
	// fmt.Println(greeting)
}
