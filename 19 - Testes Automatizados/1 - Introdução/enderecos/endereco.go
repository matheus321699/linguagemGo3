package enderecos

import "strings"

/* Função teste para validar tipo de endereço
 */
// TipoEndereco verifica se um endereco tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {

	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	// Convertendo endereco para letras minúsculas
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	
	// Pegando Primeira palavra de cada tipo de endereço
	/* 
	Função Split dentro do pacote strings que divide uma
	string em um slice de acordo com um separador passado,
	ou seja, nesse caso, de acordo com um space passado 
	como parâmetro na função Split.
	Ex: RUA DOS BOBOS, ficaria
		["RUA", "DOS", "BOBOS"]
	
	No final da função também é passado o índice ou a posição 
	do slice que se deseja pegar. Nesse caso a posição 0 foi 
	passada.
	Ex: RUA DOS BOBOS, ficaria
		["RUA"]
	*/
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	/* Verificando se a palvra passada está contida no slice
	de tipos válidos.
	*/
	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		// Deixa primeira letra de cada palavra em maiúsculo e retorna
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"

}