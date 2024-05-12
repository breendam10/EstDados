package main

import "fmt"

func calculaDistancia() {
    var distancia int

    fmt.Scanln(&distancia)
    minutos := 2*distancia
    fmt.Printf("%d minutos\n", minutos)
}



package main

import "fmt"

func calculaFolhas() {
    var jogador, qntdCompradas,folhasJogador int
    fmt.Scanln(&jogador, &qntdCompradas,&folhasJogador)

    total := jogador *folhasJogador
    if total <= qntdCompradas {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}



package main

import "fmt"

func calculaSeq() {
	var N, maxSeq, consecutivos int
	fmt.Scanln(&N)

	for i :=0; i<N; i++{
		var num int
		fmt.Scanln(&num)
		if num == 1 {
			consecutivos++
			maxSeq= max(maxSeq, consecutivos)
		} else {
			consecutivos= 0
		}
	}

	fmt.Println(maxSeq +1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}



package main

import (
	"fmt"
)

func calculaCusto() {
	var casoTeste int
	fmt.Scanln(&casoTeste)

	for c := 0; c < casoTeste; c++ {
		var pDisponiveis int
		fmt.Scanln(&pDisponiveis)

		precos := make(map[string]float64)
		for i := 0; i<pDisponiveis; i++{
			var produto string
			var preco float64
			fmt.Scan(&produto, &preco)
			precos[produto] =preco
		}
		var pRequisitados int
		fmt.Scanln(&pRequisitados)

		total :=0.0

		for i :=0; i<pRequisitados; i++{
			var produto string
			var quantidade int
			fmt.Scan(&produto, &quantidade)
			total += float64(quantidade) * precos[produto]
		}

		fmt.Printf("R$ %.2f\n", total)
	}
}




func main{
	calculaDistancia()

	calculaFolhas()

	calculaSeq()

	calculaCusto()
}
