package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("from main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("marcio.theo.92@gmail.com")
	fmt.Println(erro)
}
