package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("incrementando i")
		//time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
