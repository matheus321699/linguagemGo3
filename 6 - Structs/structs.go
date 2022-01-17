package main

import "fmt"

/* Struct: Estruturas do Go são coleções de campos digitados. 
Eles são úteis para agrupar dados em conjunto para formar registros.
*/

// Criando uma struct
type usuario struct {
	nome string 
	idade uint8
	endereco endereco
}

// Struct dentro de uma struct
type endereco struct {
	logadouro string
	numeto uint8
}

func main() {
	fmt.Println("Arquivo structs")

	// Instanciando uma struct e atribuinco valores
	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)
	// Utilizando função Printf para formatar dados
	fmt.Printf("nome = %s\n" + "2\n", u.nome)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	// Instanciando uma struct e atribuinco valores por inferência
	usuario2 := usuario{"Dave", 21, enderecoExemplo}
	fmt.Println(usuario2)

	// Passar somente um valor por inferência
	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
	

}