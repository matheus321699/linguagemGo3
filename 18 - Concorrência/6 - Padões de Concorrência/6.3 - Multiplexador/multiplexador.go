package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/* O Multiplexador vai chamar a função escrever mais de uma 
	vez, obtendo mais de um canal, e juntar esses dois canais 
	em um único canal, centralizando a comunicação em um único
	lugar dentro da função main.
	*/
	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

/* Função multiplexar que recebe dois canais como entrada e 
retorna um único canal como saída
*/

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDesaida := make(chan string)

	go func() {
		for {
			/* select que verifica quais dos dois canais tem 
			uma mensagem para ser lida. Se for no canal 1 ou 2
			ele vai jogar a mensagem para o canal de saída, ou 
			seja ele fará um encaminhamento de mensagem
			*/
			select {
			case mensagem := <-canalDeEntrada1:
				canalDesaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDesaida <- mensagem
			}
		}
		
	}()

	return canalDesaida
}

// Função escrever que retorna um canal unidirecional
func escrever(texto string) <-chan string {
	canal := make(chan string)

	// goroutine com loop infinito
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			/* time.Duration(): Calcula quanto tempo demorou para executar
			uma determinada operação e dentro dessa operação eu vou passar 
			para ele me gerar um número aleatório ou pseudoaleatório
			*/
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}