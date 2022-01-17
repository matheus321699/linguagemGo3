package main

import "fmt"

func main() {
	// Canal com sequência de números
	tarefas := make(chan int, 45)
	
	/* Canal para armazenar os resultados a medida que 
	eles forem sendo calculados
	*/
	resultados := make(chan int, 45)

	// 4 processos executando a função recursiva
	go worker(tarefas, resultados)	
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	// Enviando tarefas para o canal
	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	// Imprimindo resultado do canal resultados na tela
	for i := 0; i < 45; i++ {
		resultado := <- resultados
		fmt.Println(resultado)
	}

}

// Função worker para utilizar os canais
// Seta antes do tipo canal define um canal que só recebe dados
// Seta após o tipo canal define um canal que só envia dados
func worker(tarefas <-chan int, resultados chan<- int) {

	/* Iterando sobre o canal de tarefas pegando todos os valores
	nele e calculando o número de fibonacci e jogando o resultado no 
	canal de resultados
	*/
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

// Função recursiva de sequência de fibonacci: 1 1 2 3 5 8 13
func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
