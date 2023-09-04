package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("hello"), escrever("go dev"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalIn1, canalIn2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case message := <-canalIn1:
				canalDeSaida <- message
			case message := <-canalIn2:
				canalDeSaida <- message
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
