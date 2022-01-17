package main

import "fmt"

// Como a interface não possui nada qualquer coisa atende a interface
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	// Passando qualquer coisa para a interface
	generica("string")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "String", false, true, float64(12345))

	/* Não é sempre uma boa idéia, por questão de segurança e outros
	fatores, utilizar uma interface para passar qualquer tipo, como, por
	exemplo, no map.
	*/
	mapa := map[interface{}]interface{}{
		1:   "String",
		float32(100): true,
		"String": "String",
		true: float64(12),  
	}

	fmt.Println(mapa)
}