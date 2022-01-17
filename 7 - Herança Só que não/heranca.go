package main

import "fmt"

type pessoa struct {
	nome string 
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	/* Recurso para adicionar variáveis do tipo pessoa
	dentro do tipo ou struct estudante 
	*/
	pessoa
	curso string
	faculdade string
}

func main()  {
	fmt.Println("Herança")
	// Acesasndo variáveis dentro da struct estudante
	// pessoa2.altura = 177
	

	// Atribuindo valores as variáveis de pessoa
	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.nome)
}