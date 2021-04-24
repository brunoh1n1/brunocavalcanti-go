// Você possui uma turma com 4 alunos cujos nomes são Pedro, Ricardo, Paula e Marcia.
// Cada um fez 3 exames e tiveram notas diversas, porém a média final de cada deverá ser igual ou maior a 7.0 para que sejam aprovados.
// As notas de cada um deles são as seguintes:
// - Pedro: 4.5, 7.9 e 8.1
// - Ricardo: 5.7, 6.0 e 9.1
// - Paula: 6.8, 7.6 e 8.7
// - Marcia: 7.5, 5.9 e 7.6

// a. Crie uma variável que receba o nome de cada aluno. Estas variáveis devem possuir escopo geral.

// b. Para cada nota de cada aluno, crie uma variável. Estas variáveis de nota precisam ser com escopo de função.

// c. Em sua função main imprima:
// c.1- "Os alunos da turma são:"
// c.2- Na linha seguinte imprima o valor das variáveis com os nomes dos 4 alunos.
// c.3- Chame uma nova função chamada media

// d. a função media deverá imprimir:
// d.1- A média final de cada aluno indicando o nome do mesmo;
// d.2- Um resultado booleano que represente "true" para cada aluno que foi aprovado (cuja média seja igual ou maior a 7.0)

// e. Crie uma nova função chamada diplomas
// e.1- A função main deverá chamar também a função diplomas.
// e.2- A função diplomas irá imprimir a seguinte linha para cada alunoq ue foi aprovado:
// "Parabéns pela sua aprovação. Seu diploma estará disponível na próxima semana <NOME>."

// f. Crie uma função chamada reprovados.
// f.1- A função main também deverá chamar a função reprovados
// f.2- A função reprovados deverá imprimir o seguinte para cada aluno que foi reprovado:
// "Não sabe, não sabe, vai ter que aprender.. orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima <ALUNO>."

package main

import (
        "fmt"
)


var aluno1 = "Pedro"
var aluno2 = "Ricardo"
var aluno3 = "Paula"
var aluno4 = "Marcia"

func main() {

  fmt.Printf("\nOs alunos da turma sao: \n %v, %v, %v e %v\n", aluno1, aluno2, aluno3, aluno4)

 media();
 diplomas();
 reprovados();


}

//func media(media_01, media_a2, media_a3, media_a4) {
func media() {

  aluno1_av1 := 4.5
  aluno1_av2 := 7.9
  aluno1_av3 := 8.1
  aluno2_av1 := 5.7
  aluno2_av2 := 6.0
  aluno2_av3 := 8.7
  aluno3_av1 := 6.8
  aluno3_av2 := 7.6
  aluno3_av3 := 8.7
  aluno4_av1 := 7.5
  aluno4_av2 := 5.9
  aluno4_av3 := 7.9

  media_a1 :=  (aluno1_av1 + aluno1_av2 + aluno1_av3) / 3
  media_a2 :=  (aluno2_av1 + aluno2_av2 + aluno2_av3) / 3
  media_a3 :=  (aluno3_av1 + aluno3_av2 + aluno3_av3) / 3
  media_a4 :=  (aluno4_av1 + aluno4_av2 + aluno4_av3) / 3

  
  fmt.Printf("\nA media de %v e: %.2f - situacao de aprovado: %v\n", aluno1, media_a1, media_a1 >= 7)
  fmt.Printf("A media de %v e: %.2f - situacao de aprovado: %v\n", aluno2, media_a2, media_a2 >= 7)
  fmt.Printf("A media de %v e: %.2f - situacao de aprovado: %v\n", aluno3, media_a3, media_a3 >= 7)
  fmt.Printf("A media de %v e: %.2f - situacao de aprovado : %v\n", aluno4 , media_a4, media_a4 >= 7)
  }

  func diplomas(){
          
    aluno1_av1 := 4.5
    aluno1_av2 := 7.9
    aluno1_av3 := 8.1
    aluno2_av1 := 5.7
    aluno2_av2 := 6.0
    aluno2_av3 := 8.7
    aluno3_av1 := 6.8
    aluno3_av2 := 7.6
    aluno3_av3 := 8.7
    aluno4_av1 := 7.5
    aluno4_av2 := 5.9
    aluno4_av3 := 7.9
  
    media_a1 :=  (aluno1_av1 + aluno1_av2 + aluno1_av3) / 3
    media_a2 :=  (aluno2_av1 + aluno2_av2 + aluno2_av3) / 3
    media_a3 :=  (aluno3_av1 + aluno3_av2 + aluno3_av3) / 3
    media_a4 :=  (aluno4_av1 + aluno4_av2 + aluno4_av3) / 3       
          
    fmt.Printf("\n\n")
    if media_a1 >= 7 {
      fmt.Println("Parabéns pela sua aprovação. Seu diploma estará disponível na próxima semana .", aluno1)
    }
    if media_a2 >= 7 {
      fmt.Println("Parabéns pela sua aprovação. Seu diploma estará disponível na próxima semana .", aluno2)
    }
    if media_a3 >= 7 {
      fmt.Println("Parabéns pela sua aprovação. Seu diploma estará disponível na próxima semana .", aluno3)
    }
    if media_a4 >= 7 {
      fmt.Println("Parabéns pela sua aprovação. Seu diploma estará disponível na próxima semana .", aluno4)
    }
          
  }

  func reprovados(){
          
    aluno1_av1 := 4.5
    aluno1_av2 := 7.9
    aluno1_av3 := 8.1
    aluno2_av1 := 5.7
    aluno2_av2 := 6.0
    aluno2_av3 := 8.7
    aluno3_av1 := 6.8
    aluno3_av2 := 7.6
    aluno3_av3 := 8.7
    aluno4_av1 := 7.5
    aluno4_av2 := 5.9
    aluno4_av3 := 7.9
  
    media_a1 :=  (aluno1_av1 + aluno1_av2 + aluno1_av3) / 3
    media_a2 :=  (aluno2_av1 + aluno2_av2 + aluno2_av3) / 3
    media_a3 :=  (aluno3_av1 + aluno3_av2 + aluno3_av3) / 3
    media_a4 :=  (aluno4_av1 + aluno4_av2 + aluno4_av3) / 3       
          
    fmt.Printf("\n\n")
    if media_a1 < 7 {
      fmt.Println("Não sabe, não sabe, vai ter que aprender.. orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima ", aluno1)
    }
    if media_a2 < 7 {
      fmt.Println("Não sabe, não sabe, vai ter que aprender.. orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima ", aluno2)
    }
    if media_a3 < 7 {
      fmt.Println("Não sabe, não sabe, vai ter que aprender.. orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima ", aluno3)
    }
    if media_a4 < 7 {
      fmt.Println("Não sabe, não sabe, vai ter que aprender.. orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima ", aluno4)
    }
          
  }



