package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)

	go func() {
		write("hello world")
		waitGroup.Done()
	}()

	go func() {
		write("programando em go!")
		waitGroup.Done()
	}()

	go func() {
		write("routine 3!")
		waitGroup.Done()
	}()

	go func() {
		write("routine 4!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
