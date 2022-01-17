package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// Desclarando um Array
	var array1 [5]string
	// Atribuindo valores ao array
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	/* Nessa forma de declarar um array os três pontos vão fixar o tamanho 
	do array de acordo com a quantidade de valores passados
	*/
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	fmt.Println(array3[1])

	// Declarando um slice
	slice := []int{10, 11, 12, 14, 15, 16, 17}
	fmt.Println(slice)

	// Type: função que retorna o tipo do objeto
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))


	/* No slice existe uma função append que adiciona um item no slice
	e retornar um slice novo com esse item adicionado
	*/
	slice = append(slice, 18)
	fmt.Println(slice)
	fmt.Println(slice[1])

	// Pegando um pedaço do array e salvando dentro do slice
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	/* Arrays Internos
	*/
	// Utilizando função make utilizada para alocar espaço na memória para uma estrutura
	fmt.Println("--------------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len((slice3))) // length
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len((slice3))) // length
	fmt.Println(cap(slice3)) // capacidade
	
}