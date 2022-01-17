package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	/* Primeiro tipo de for
	*/
	for i < 10 {
		// A biblioteca time é utilizada para manipular o tempo
		// A função Sleep pausa o cógigo
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)	
	}

	fmt.Println(i)

	/* Segundo tipo de for
	*/
	for j := 0; j < 10; j+= 2 {
		fmt.Printf("Incremento %d\n", j)
		time.Sleep(time.Second)
	}

	/* Terceiro tipo de for com a cláusula range, utilizado
	para iterar sobre algo, como, por exemplo, em um array ou slice.
	*/
	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// Iterando sobre uma string
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	// Iterando por um map
	usuario := map[string]string {
		"nome": "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Obs: Não existe a possibilidade de iterar sobre uma struct

	// Loop infinito
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

}
