package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func write(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	fmt.Println(totalDaSoma)

	write("numeros", 1, 2, 3, 4, 5, 6, 7)
}
