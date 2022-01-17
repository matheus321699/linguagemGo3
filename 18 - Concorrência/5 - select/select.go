package main

import (
	"fmt"
	"time"
)

func main() {
	/* Duas goroutines rodando ao mesmo tempo que 
	vão enviar valores para canais diferentes
	*/

	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Microsecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {

		/* Estrutura Condicional para goroutines que não necessitam esperar
		determinado tempo para realizar uma interação com o canal1
		*/
		select {
			// Caso canal1 está pronto para receber um valor printa o valor na tela
			case mensagemCanal1 := <-canal1:
				fmt.Println(mensagemCanal1)
			
			// Caso canal2 está pronto para receber um valor printa o valor na tela
			case mensagemCanal2 := <-canal2:
				fmt.Println(mensagemCanal2)
		}

//		mensagemCanal1 := <-canal1
		// Mensagem chega do canal1 e espera 0,5 segundo para ir para o canal2
//		fmt.Println(mensagemCanal1)

//		mensagemCanal2 := <-canal2
		// Mensagem chega do canal2 e espera 2 segundo para ir para o canal1
//		fmt.Println(mensagemCanal2)

		/* Obs: perda de tempo na execução das funções
		*/

	}

}