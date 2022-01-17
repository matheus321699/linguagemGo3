package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	/* Criando um if init que executa uma condição if e inicializa
	uma variável já nessa condição. Nesse caso a variável fica limitada 
	ao escopo do if não podendo ser utilizada fora dele.
	*/
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que 10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}