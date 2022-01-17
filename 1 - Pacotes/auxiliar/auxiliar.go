package auxiliar

import "fmt"

/* Para uma função ser pública e ser vista por qualauer
outro pacote o seu nome deve iniciado com letra minúscula,
e para ser uma função que só pode ser vista dentro do seu
pacote.
*/
// Função que registra uma mensagem na tela
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
