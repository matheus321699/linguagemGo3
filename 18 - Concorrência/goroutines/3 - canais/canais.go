package main

import (
	"fmt"
	"time"
)

func main() {
	//Criando e atribuindo um tipo ao canal 
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever a ser ")

	// Canal vai esperar receber um valor
	for {
	// variável aberto de Retorno para verificar se o canal está aberto
	// Canal espera um valor e salva na variável mensagem
	mensagem, aberto := <-canal // Canal interrompe o fluxo de execução após receber a mensagem
	
		// Se o canal não estiver aberto sai do for
		if !aberto {
			break
		}

	fmt.Println(mensagem)
	}

	// Outra forma de deixar o canal aberto
	/* Para cada mensagem que for recebida no canal enquanto
	ele estiver aberto irá printar na tela.
	*/
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

// Passando canal por parâmetro
func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// Enviar dado por um canal
		canal <- texto
		time.Sleep(time.Second)
	}

	// Fechando canal após sair do for
	close(canal)
}

