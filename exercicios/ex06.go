// 1. Crie um programa que, no escopo de pacote crie as seguintes variáveis:
// - nome (string)
// - idade (int)
// - peso (float64)
// - pesopizza (int)
// 
// 	a. para nome, atribua seu nome
// 	b. para idade, atribua sua idade
// 	c. para peso, atribua seu peso
// 
// 2. Na função main:
// 
// Imprima o seguinte parágrafo da forma que você desejar,
// respeitando as seguintes regras:
// 
// 	a. Onde você encontrar <nome>,
// 	deverá ser impresso o valor da variável nome. Onde encontrar <idade>, deverá
// 	ser impresso o valor da variável idade. E assim por diante.
// 	b. Onde você encontrar <tipo_variável_nome>, deverá ser impresso o tipo da variável
// 	nome e assim por diante.
// 
// 
// Meu nome é <nome>. Estou aprendendo o básico de Go e git.
// Possuo <idade> anos e estou pesando aproximadamente <peso>.
// Utilizei variáveis para lhe passar estas informações pessoais com Go.
// Por exemplo, utilizei a variável nome, com o tipo <tipo_variável_nome>.
// Utilizei outras duas variáveis, para idade e peso, com os tipos <tipo_variável_idade> e <tipo_variável_peso>.
// Gosto muito de pizza, portanto fui em um rodízio de pizzas ontem e acabei engordando 1 kg.
// 
// 3. Após imprimir o texto acima, atribua o valor da variável pesopizza de forma que ele seja o seu peso + 1kg.
// Em seguida imprima o valor e o tipo da variável pesopizza

package main

import (
	"fmt"
)

var nome = "Bruno"
var idade = 29
var peso = 99.5
var pesopizza int

func main() {
	//fmt.Println(x)
	//fmt.Printf("%v %v %v", x, y, z)
	//fmt.Scanf("%v", &s)
	fmt.Printf("Meu nome é %v. Estou aprendendo o básico de Go e git.\nPossuo %v anos e estou pesando aproximadamente %v.\nUtilizei variáveis para lhe passar estas informações pessoais com Go.\nPor exemplo, utilizei a variável nome, com o tipo %T.\nUtilizei outras duas variáveis, para idade e peso, com os tipos %T e %T.\nGosto muito de pizza, portanto fui em um rodízio de pizzas ontem e acabei engordando 1 kg.\n", nome, idade, peso, nome, idade, peso)
    pesopizza = int(peso) + 1.0
	fmt.Printf("peso atia após rodizio: %v\ntipo variavel pesopizza %T\n", pesopizza, pesopizza)
}