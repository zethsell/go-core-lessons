package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	number := 10

	if number > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("menor ou igual a 15")
	}

	if outroNumero := number; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if outroNumero < -10 {
		fmt.Println("Numero é menor que 10")
	} else {
		fmt.Println("Entre 0 e 10")
	}

}
