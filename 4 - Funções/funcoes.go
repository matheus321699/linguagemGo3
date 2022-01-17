package main

import "fmt"

/* Declarando função fora da função main
*/
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

/* Função com mais de um retorno:
*/
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	/* Declarando variável do tipo função e inserir uma função dentro dela
	*/
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	/* O uso do _(underline) na declaração de uma variável ignora a declaração
	da variável para que não seja necessário a sua utilização
	*/
	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)
}