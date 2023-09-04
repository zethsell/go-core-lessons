package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Numero Invalido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Numero Invalido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(3)
	fmt.Println(dia)

	dia = diaDaSemana2(5)
	fmt.Println(dia)
}
