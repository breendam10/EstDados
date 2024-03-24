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



  var numCasos int
  fmt.Scan(&numCasos)
  for i := 0; i < numCasos; i++ {
	  var num1, den1, num2, den2 int
	  var operador string
	  fmt.Scan(&num1, &operador, &den1, &num2, &den2)
	  numResultado, denResultado := realizarOperacao(num1, den1, operador, num2, den2)
	  numSimplificado, denSimplificado := simplificarFracao(numResultado, denResultado)
	  fmt.Printf("%d/%d = %d/%d\n", numResultado, denResultado, numSimplificado, denSimplificado)
 		 }
	}




func mdc(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func simplificarFracao(num, den int) (int, int) {
    maiorDivisorComum := mdc(num, den)
    numSimplificado := num / maiorDivisorComum
    denSimplificado := den / maiorDivisorComum
    return numSimplificado, denSimplificado
}

func realizarOperacao(num1, den1 int, operador string, num2, den2 int) (int, int) {
    var numResultado, denResultado int
    switch operador {
    case "+":
        numResultado = num1*den2 + num2*den1
        denResultado = den1 * den2
    case "-":
        numResultado = num1*den2 - num2*den1
        denResultado = den1 * den2
    case "*":
        numResultado = num1 * num2
        denResultado = den1 * den2
    case "/":
        numResultado = num1 * den2
        denResultado = num2 * den1
    }
    return numResultado, denResultado





	var N int
	var L, Q float64


	fmt.Scan(&N, &L, &Q)


	aguaTotal := L


	participantes := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&participantes[i])
	}
	aguaPorParticipante := Q
	lastParticipante := ""
	for _, participante := range participantes {

		if aguaTotal <= aguaPorParticipante {
			lastParticipante = participante
		}
	}
	
	fmt.Printf("%s %.1f\n", lastParticipante, Q)
}






