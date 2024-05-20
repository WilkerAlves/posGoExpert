package main

import "fmt"

// está-se criando um constante de nome constHelloWorld e atribuindo o valor "hello, world"
// Go possui inferencia de tipos, isso significa que como estamos passando o valor "hello, world" que é um texto, logo a constante constHelloWorld é do tipo string
const constHelloWorld = "hello, world"

// variaveis declaradas forá de funções possuem o seu escopo global. Isso significa que todas as funções podem acessa-las e alterar seus valores
var (
	// está-se criando um variavel do tipo boleano(verdadeiro/falso)
	b bool
	// está-se criando um variavel do tipo inteiro(1_pacoteEFuncMain,2_declaracaoEAtribuicaoVariaveEConstante,3,4...)
	c int
	// está-se criando um variavel do tipo texto(olá mundo)
	d string
	// está-se criando um variavel do tipo ponto flutuante(1_pacoteEFuncMain.5,5.9,6.05,...)
	e float64
)

// Em go tambem é possivel declarar variavies a atribuir valor
var (
	// está-se criando um variavel do tipo boleano(verdadeiro/falso) e atribuindo o valor true
	t bool = true
	// está-se criando um variavel do tipo inteiro(1_pacoteEFuncMain,2_declaracaoEAtribuicaoVariaveEConstante,3,4...) e atribuindo o valor 1_pacoteEFuncMain
	v int = 1
	// está-se criando um variavel do tipo texto(olá mundo) e atribuindo o valor de teste
	n string = "teste"
	// está-se criando um variavel do tipo ponto flutuante(1_pacoteEFuncMain.5,5.9,6.05,...) e atribuindo o valor de 1_pacoteEFuncMain.5
	m float64 = 1.5
)

func main() {
	// variaveis declaradas dentro de um função, tem o seu escopo local. Isso significa que elas podem ser acessadas e alteradas apenas dentro da função que foram declaradas
	var (
		x bool
		y int
		z string
		g float64 = 5.6
	)

	// atribuindo valor a uma variavel
	z = "atribuindo valor"
	fmt.Println(z)

	// apos atribuir um valor a uma constante não é possivel altera-lo
	//constHelloWorld = "teste"
	fmt.Println(constHelloWorld)

	// valores default das variaveis
	// globais
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	// locais
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	// declarando e atribuindo valor as variavies
	fmt.Println(t)
	fmt.Println(v)
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(g)

	// outra forma de criar e atribuir valor a variaveis
	segundaString := "test"
	segundaInt := 20
	segundaBool := true
	segundaFloat := 5.6

	fmt.Println(segundaString)
	fmt.Println(segundaInt)
	fmt.Println(segundaBool)
	fmt.Println(segundaFloat)

}
