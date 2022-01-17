package main

import "fmt"

func main() {
	// Criando canal e especificando capacidade para o canal
	// Canais com buffer não interrompem a execução
	canal := make(chan string, 2)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <- canal
	mensagem2 := <- canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}