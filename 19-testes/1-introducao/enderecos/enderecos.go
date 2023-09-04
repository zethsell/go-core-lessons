package enderecos

import (
	"strings"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoLowerCase := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoLowerCase, " ")[0]

	tipoEnderecoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			tipoEnderecoValido = true
		}
	}

	if tipoEnderecoValido {
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo Inválido"

}
