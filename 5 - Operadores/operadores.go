package main

import "fmt"

func main() {
	/* ARITMÉTICOS
	*/
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// Em somente variáveis do mesmo tipo podem realizar operações aritméticas
	var numero1 int16 = 10
	var numero2 int32 = 25
	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)
	// Fim dos aritméticos

	/* ATRIBUIÇÃO
	*/
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// Fim dos operadores de atribuição

	/* OPERADORES RELACIONAIS
	*/
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// Fim dos relacionais

	/* OPERADORES LÓGICOS
	*/
	fmt.Println("--------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(falso)
	// Fim dos operadores lógicos

	/* OPERADORES UNÀRIOS
	*/
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero-=20 // numero = numero - 20
	fmt.Println(numero)

	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)
	// Fim dos operadores unários

	/* Em Go não existe operador ternário, o que pode ser resolvido
	utilizando if e else
	*/
	var texto string 
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}