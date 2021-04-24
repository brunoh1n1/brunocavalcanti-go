// Utilizando o código do exercício 04...
// 1. No escopo de pacote, atribua os seguintes valores
// para as 3 variáveis:
//
// a. para x, atribua 42
// b. para y, atribua "James Bond"
// c. para z, atribua true
//
// 2. Na função main:
//
// a. utilize fmt.Sprintf para imprimir todos os VALORES em
// uma única string. ATRIBUA o valor retornado do TIPO string
// utilizando o declaradaor curto (short) para uma VARIÁVEL
// com o IDENTIFICADOR "s".
// b. imprima o valor de "s"

package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true
var s string

func main() {
	fmt.Println(y)
	fmt.Printf("%v %v %v", x, y, z)
	fmt.Scanf("%v", &s)
	fmt.Printf("\n%v\n", s)
}
