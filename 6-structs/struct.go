package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	enderecoEx := endereco{"rua 1", 1}

	var user usuario
	user.nome = "Marcio"
	user.idade = 31
	fmt.Println(user)

	user2 := usuario{"Marcio", 31, enderecoEx}
	fmt.Println(user2)

	user3 := usuario{idade: 21}
	fmt.Println(user3)
}
