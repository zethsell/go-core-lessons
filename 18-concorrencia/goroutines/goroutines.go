package main

import (
	"fmt"
	"time"
)

func main() {
	go write("hello world")
	write("programando em go!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
