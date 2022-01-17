package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// Sincronizando função escrever ou goroutines com wait group

	// Declarando waitGroup
	var waitGroup sync.WaitGroup

	// Definindo quantidade de goroutines que vão estar executando ao mesmo tempo
	// e fazem parte do grupo de espera
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo!")
		// Retira uma goroutine do contador
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go!")
		// Retira uma goroutine do contador
		waitGroup.Done()
	}()

	/* Diz para o programa ou a função main esperar
	a contagem das goroutimes chegar a 0.
	*/
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
