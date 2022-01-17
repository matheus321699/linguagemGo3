// TESTE DE UNIDADE 

package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

/* Refatorando função TestTipoEndereco para que tenha várias 
das structs criada:
*/
// Função que vai testar a função tipoDeEndereco
// Assinatura de uma função de teste
func TestTipoDeEndereco(t *testing.T) {
	
	cenariosDeTeste := []cenarioDeTeste {
		{ "Rua ABC", "Rua"},
		{ "Avenida Paulista", "Avenida"},
		{ "Rodovia dos Imigrantes", "Rodovia"},
		{ "Praça das Rosas", "Tipo Inválido"},
		{ "Estrada Qualquer", "Estrada"},
		{ "RUA DOS BOBOS", "Rua"},
		{ "AVENIDA REBOUÇAS", "Avenida"},
		{ "", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebidp %s é diferente do esperado %s", 
			retornoRecebido,
			cenario.retornoEsperado,
		)

		}
	}
}