package main

import "fmt"

func main() {

	// Declarando map e atribuindo valores
	usuario := map[string]string {
		"nome" : "Pedro",
		"sobrenome" : "Silva",
	}

	// Acessando chave nome
	fmt.Println(usuario["nome"])

	// Declarando map alinhado e atribuindo valores
	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "João",
			"ultimo": "Pedro",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus 1",
		},

	}
	
	fmt.Println(usuario2)

	// Apagando uma chave do map utilizando uma função nativa do Go chamada delete
	delete(usuario2, "nome")
	fmt.Println(usuario2)
	
	// Adicionando uma chave no map 
	usuario2["signo"] = map[string] string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)

}