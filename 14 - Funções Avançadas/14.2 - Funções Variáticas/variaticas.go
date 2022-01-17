package main

import "fmt"

// Criando uma função variática, que vai receber de 0 a n int parâmetros
/* Obs: Só é possível ter um parâmetro variático por função e obrigatoriamente
o parâmetro variático tem que ser o último
*/
func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

/* Combinando uma função que recebe parâmetros fixos
com parâmetros variáticos
*/ 
func escrever(texto string, numeros ...int){
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}


func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 20, 10, 12, 13)
	fmt.Println(totalDaSoma)

	escrever("Olá mundo", 10, 2, 3, 4, 5, 6, 7)
}