package main

import (
	"errors"
	"fmt"
)

func main() {
	/* Tipos de inteiros em Go e a quantidade
	de bits que eles ocupam.
	*/
	// int8, int16, int32, int64
	
	var numero int16 = 10
	fmt.Println(numero)

	/* Por padrão caso não seja especificado 
	o número de bits o Go adota por padraão o número
	de bits da aquitetura de seu computador.
	*/
	var numero2 int = -100
	fmt.Println(numero2)

	numero3 := 10000
	fmt.Println(numero3)

	// unit -> unsygned int ou int sem sinal
	var numero4 uint32 = 1000
	fmt.Println(numero4)

	// alias
	// INT32 = RUNE
	var numero5 rune = 1234
	fmt.Println(numero5)

	// BYTE = UINT8
	var numero6 byte = 123
	fmt.Println(numero6)

	/* Valores reias em Go:
	*/
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000.45
	fmt.Println(numeroReal2)

	// Inferência de tipos valores reais
	numeroReal3 := 12345.67 
	fmt.Println(numeroReal3)

	/* Valores do tipo String
	*/
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// Número do caracter na tabela ASK
	char := 'B'
	fmt.Println(char)

	// Valor zero ou valor que é atribuído automaticamente a variável
	var texto string
	fmt.Println(texto)

	/* Valores Booleanos
	*/
	var booleano1 bool
	fmt.Println(booleano1)

	/* Tipo de dado Erro
	*/
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}

