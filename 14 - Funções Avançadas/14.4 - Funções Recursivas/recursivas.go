package main

import "fmt"

// Função recursiva de sequência de fibonacci: 1 1 2 3 5 8 13
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	posicao := uint(15)

	// Retornando todos os valores da sequência de fibonacci
	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}


	fmt.Println(fibonacci(posicao))
}