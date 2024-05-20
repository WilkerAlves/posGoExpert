package main

import "fmt"

// criando um tipo, isso é muito usado para deixar o codigo mais claro
// O novo tipo criado usa os tipos base do go
// essa estrategia é muito usada para aumenta a expresividade
type MeuTipo int

var (
	// está-se criando um variavel do tipo inteiro(1_pacoteEFuncMain,2_declaracaoEAtribuicaoVariaveEConstante,3,4...)
	c int

	f MeuTipo = 6
)

func main() {
	fmt.Println(f)

}
