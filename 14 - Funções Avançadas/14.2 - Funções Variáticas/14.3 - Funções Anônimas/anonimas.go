package main

import "fmt"

func main() {

	// Declara função anônima
	func (texto string) {
		fmt.Println(texto)
	}("Passando Parâmetro") // Executa função anônima

	// Declarando função anônima e armazanando retorno na variável
	retorno := func (texto string) string{
		// Função Sprintf concatena informações
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro") // Executa função anônima

		fmt.Println(retorno)

}