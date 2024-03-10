
package main


import "fmt"

func vetorint() {
	rand.Seed(time.Now().UnixNano())

	vetor := [10]int{}

	for i := 0; i < 10; i++ {
		vetor[i] = rand.Intn(100) 
	}
	for i := 0; i < 10; i++ {
		fmt.Println(vetor[i])
	}
}



package main

import (
	"fmt"
)

func main() {

	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)


	reversed := convertString(input)

	
	fmt.Println("String invertida:", reversed)
}

func convertString(s string) string {

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}



package main

import "fmt"

type Ponto struct {
	X float64
	Y float64
}

func (p Ponto) DistanciaOrigem() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func calculaDistancia() {

	p := Ponto{X: 3, Y: 4}


	dist := p.DistanciaOrigem()

	fmt.Printf("A distância do ponto (%.2f, %.2f) até a origem é %.2f\n", p.X, p.Y, dist)
}





