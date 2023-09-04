package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "Hello"
	channel <- "Hello2"

	message := <-channel
	fmt.Println(message)

	message2 := <-channel
	fmt.Println(message2)
}
