package main

import "fmt"

func main() {
	task := make(chan int, 45)
	result := make(chan int, 45)

	go worker(task, result)
	go worker(task, result)
	go worker(task, result)
	go worker(task, result)
	go worker(task, result)
	go worker(task, result)
	go worker(task, result)
	go worker(task, result)

	for i := 0; i < 45; i++ {
		task <- i
	}

	close(task)

	for i := 0; i < 45; i++ {
		resultado := <-result
		fmt.Println(resultado)
	}
}

func worker(task <-chan int, result chan<- int) {
	for numero := range task {
		result <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
