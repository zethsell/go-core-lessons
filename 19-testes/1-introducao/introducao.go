package main

import (
	"fmt"
	"introducao/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua das rosas")
	fmt.Println(tipoEndereco)
}
