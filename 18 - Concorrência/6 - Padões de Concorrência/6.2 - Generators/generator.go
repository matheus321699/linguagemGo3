package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo!")

	// Interagindo com a função escrever com for
	for i := 0; i < 10; i++ {
		fmt.Print(<-canal)
	}

}

// Função escrever que retorna um canal unidirecional
func escrever(texto string) <-chan string {
	canal := make(chan string)

	// goroutine com loop infinito
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}