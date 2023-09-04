package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("hello world", channel)

	fmt.Println("post execution")

	//for {
	//	message, open := <-channel
	//	if !open {
	//		break
	//	}
	//	fmt.Println(message)
	//}

	for message := range channel {
		fmt.Println(message)
	}
}

func write(text string, channel chan string) {
	//time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
