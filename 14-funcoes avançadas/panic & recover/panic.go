package main

import "fmt"

func funcao1() {
	fmt.Println("exec func 1")
}

func funcao2() {
	fmt.Println("exec func 2")
}

func recoveryProcess() {
	if r := recover(); r != nil {
		fmt.Println("Execuçao recuperada")
	}
	println("tentativa de recuperar!")
}

func alunoAprovado(n1, n2 float32) bool {
	defer recoveryProcess()
	defer fmt.Println("Média calculada. Resultado será retornado!")
	media := (n1 + n2) / 2

	if media == 6 {
		panic("A MEDIA É EXATAMENTE 6!")
	}

	return media > 6
}

func main() {
	//defer funcao1()
	//funcao2()
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("pos execucao")
}
