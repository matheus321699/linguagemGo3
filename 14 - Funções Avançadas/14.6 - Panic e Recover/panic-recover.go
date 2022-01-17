package main

import "fmt"

func recuperarExecucao() {
	// Utilizando a clausula recover
	// Caso a função não estiver entrando em panic o recover é ignorado
	if r := recover(); r!= nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	// A média do aluno no programa nunca pode ser 6
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	
	panic("A MÉDIA È EXATAMENTE 6!")
}

func main() {
fmt.Println(alunoEstaAprovado(7, 6))
fmt.Println("Pós execução")

}

