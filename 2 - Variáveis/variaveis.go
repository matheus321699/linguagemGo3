package main

import "fmt"

func main() {
	var variavel string = "Variável 1"

	/* Declarando variável por inferência de tipo,
	onde a variável é determinada pelo valor dela
	*/
	variavel2 := "Variável 2"
	fmt.Println(variavel)
	fmt.Println(variavel2)

	/* Declarando mais de uma variável de uma vez
	*/
	var (
		variavel3 string = "lala"
		variavel4 string = "lalala"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	/* Declaração de constante
	*/
	const constante string = "Constante 1"
	fmt.Println(constante)

	/* Trocar valor de duas variáveis
	*/
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
