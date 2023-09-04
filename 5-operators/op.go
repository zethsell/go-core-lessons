package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	mod := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, mod)

	var var1 string = "string"
	var2 := "string2"
	fmt.Println(var1, var2)

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	fmt.Println("---------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	fmt.Println("---------")
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20

	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)
	fmt.Println("---------")

}
