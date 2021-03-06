package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"

	default: 
		return "Número Inválido"
	}
}

/* Outra maneira de avaliar as condições do switch, onde pode ser
utilizada mais de uma variável para verificar um condição.
*/
func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		/* A cláusula fallthrough avalia a condição e executa
		o case e logo após vai jogar o seu código para dentro da 
		próxima condição sem verificar o case.
		*/
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número Inválido"
	}

	return diaDaSemana
}


func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(7)
	fmt.Println(dia)

	fmt.Println("---------------------")
	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}