/*
- Na primeira linha de qualquer arquivo go sempre terá o pacote
- O nome do pacote é nome do diretorio que o arquivo está
- O pacote main é um caso especial, isso devido a ser o pacote principal da linguagem, sendo o ponto de entrada da aplicação
- Todas as funções variaveis(declaradas fora de funções), constantes... estaram disponiveis para todos os arquivos que fazem parte do pacote
- Isso acontece pois go entender que tudo o que está dentro de um pacote é um unico arquivo
*/
package main

func main() {
	print("hello world")
}
