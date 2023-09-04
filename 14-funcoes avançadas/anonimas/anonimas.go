package main

import "fmt"

func main() {
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("parametro")

	fmt.Println(retorno)
}
