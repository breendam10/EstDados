package main

import (
  "fmt"
)

func main() {
  var Peca1, qPecas1 int
  var valorPeca1 float64
  var Peca2, qPecas2 int
  var valorPeca2 float64

  fmt.Scan(&Peca1, &qPecas1, &valorPeca1)
  fmt.Scan(&Peca2, &qPecas2, &valorPeca2)
  Total := (float64(qPecas1) * valorPeca1) + (float64(qPecas2) * valorPeca2)

  fmt.Printf("Valor a pagar: R$ %.2f\n", Total)
}
