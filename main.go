package main

import (
	"fmt"
)

func main() {
  var twoNumbersOperation[4]int{0, 1, 2, 3}
  var operation int
	fmt.Println("CALCULADORA:")
  fmt.Println("Digite o numero da operacao desjada:")
  fmt.Println("1 - SOMA")
  fmt.Println("2 - SUBTRACAO")
  fmt.Println("3 - MULTIPLICACAO")
  fmt.Println("4 - DIVISAO")
  fmt.Println("5 - ELEVAR NUMERO AO QUADRADO")
  fmt.Println("6 - ELEVAR NUMERO AO CUBO")
  fmt.Println("7 - RAIZ QUADRADA")
  fmt.Println("8 - RAIZ CUBICA")
  fmt.Println("9 - LOGARITMO")
  fmt.Scanln(&operation)
  for i := range twoNumbersOperation {
    if twoNumbersOperation[i].Key == operation {
      
    }
  }
}
