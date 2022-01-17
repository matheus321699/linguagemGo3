package main

import (
	"fmt"
	"math"
)

/* Criando uma interface para utilizar função de escrever a área do tipo
triangulo e retangulo. As interfaces só ppssuem assinaturas de métodos,
que dizem como os métodos devem ser.
*/
type forma interface {
	area() float64
}


func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

/* Para que o retangulo satisfaça uma interface e possa ser passado como
parâmetro é necessário que a struct do retângulo tenha um método com a funçao
area igual ao da interface.
*/
type retangulo struct {
	altura float64
	largura float64
}

// Criando um método que possui uma função area() para o retangulo
func (r retangulo) area() float64 {
	return r.altura * r.altura
}

/* Para que o circulo satisfaça uma interface e possa ser passado como
parâmetro é necessário que a struct do circulo tenha o método
area igual ao da interface.
*/
type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}