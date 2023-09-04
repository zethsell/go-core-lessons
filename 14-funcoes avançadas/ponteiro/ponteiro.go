package main

import "fmt"

func invertSinal(numero int) int {
	return numero * -1
}

func invertSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := invertSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	invertSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
