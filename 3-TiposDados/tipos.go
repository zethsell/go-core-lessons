package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int64 = 1001233123
	fmt.Println(number)

	var number2 uint64 = 1001233123
	fmt.Println(number2)

	//alias
	//int32 = RUNE
	var number3 rune = 12345
	fmt.Println(number3)

	//uint8 = BYTE
	var number4 byte = 123
	fmt.Println(number4)

	var number5 float32 = 123.45
	fmt.Println(number5)

	number6 := 12345.67
	fmt.Println(number6)

	var str string = "texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//retorna codigo da tabela ASCII
	char := 'B'
	fmt.Println(char)

	var text int16
	fmt.Println(text)

	var boolean1 bool
	fmt.Println(boolean1)

	var err error = errors.New("Erro Interno")
	fmt.Println(err)
}
