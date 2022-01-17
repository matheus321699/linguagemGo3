package main

import "fmt"

// Função que inverse sinal de um número utilizando uma cópia da variável passada
func inverterSinal(numero int) int {
	return numero * -1
}

// Função que inverse sinal de um número utilizando o endereço de memória
// da variável passada
func inverterSinalComPonteiro(numero *int) {
	// Desreferênciação da variável
	*numero = *numero * -1
}

func main() {

	numero := 20
	// Passando parâmetro por cópia
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	// Passando endereço de memória para a função
	// Passando parâmetro por referência
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}