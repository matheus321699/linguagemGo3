package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

/* Pacote main que vai importar a aplicação de outro
pacote criado onde vai ser gerado a aplicação com os
comando criados.
*/
func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	
	// Tratando erro
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}