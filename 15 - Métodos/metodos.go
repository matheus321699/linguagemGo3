package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

// Declarando método
// Todos os usuários tem um método chamado salvar
// Método de usuario
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// Método que altera algum valor de variável do usuário
func (u *usuario) fazerAniversario() {
	u.idade++
}


func main() {
	usuario1 := usuario{"Usuário 1", 20}
	// Não precisa passar o usuário como parâmetro
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}