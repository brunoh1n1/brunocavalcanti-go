// 1. Utilizando o operador para declaração curta (ou short) de variável,
// ATRIBUA os seguintes VALORES para os IDENTIFICADORES "x", "y" and "z":

// a. 42
// b. "James Bond"
// c. true

// 2. Agora imprima os valores armazenados nestas variáveis
// utilizando:
// a. Uma única chamada de Print (Um único Print)
// b. Múltiplas chamadas de Print.

package main

import (
  "fmt"
)

func main(){
  x := 42
  y := "James Bond"
  z := true

  fmt.Println("Resposta a: Um unico print")
  fmt.Printf("%v\n%v\n%v", x, y, z)
  fmt.Printf("\nResposta b:  vários prints\n")
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)
}
