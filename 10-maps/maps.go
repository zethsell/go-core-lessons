package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Marcio",
			"segundo":  "Theodoro",
		},
		"curso": {
			"primeiro": "Marcio",
			"segundo":  "Theodoro",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "curso")
	fmt.Println(usuario2)
}
