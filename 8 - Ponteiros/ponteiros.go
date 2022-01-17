package main

import "fmt"

func main() {
	
	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	fmt.Println("Ponteiros")

	// Atribuindo valores por cópia
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)
	
	// Atribuindo valores por ponteiros
	var variavel3 int = 100
	// Criando um ponteiro
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)

	/* Ver valor da variável através do ponteiro que está salvo 
	na memória
	*/
	fmt.Println(variavel3, *ponteiro) // desreferenciação

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)
}