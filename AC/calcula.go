package main

import (
	"fmt"
)

func main() {
	var quantidade int
	fmt.Print("Digite a quantidade de números: ")
	fmt.Scanln(&quantidade)

	valores := make([]float64, quantidade)
	for i := 0; i < quantidade; i++ {
		fmt.Printf("Digite o número %d: ", i+1)
		fmt.Scanln(&valores[i])
	}

	media := calculaMedia(valores...)
	fmt.Println("A média dos números é:", media)
}

func calculaMedia(valores ...float64) float64 {
	total := 0.0
	for _, valor := range valores {
		total += valor
	}
	return total / float64(len(valores))
}


func main() {
    var celsius float64
    fmt.Print("Digite a temperatura em Celsius: ")
    fmt.Scanln(&celsius)

    fahrenheit := (celsius * 9 / 5) + 32
    fmt.Println("A temperatura em Fahrenheit é:", fahrenheit)
}

func main() {
    var base, expoente int
    fmt.Print("Digite a base: ")
    fmt.Scanln(&base)
    fmt.Print("Digite o expoente: ")
    fmt.Scanln(&expoente)

    resultado := calcularPotencia(base, expoente)
    fmt.Println(resultado)
}

func calcularPotencia(base, expoente int) int {
    resultado := 1
    for i := 0; i < expoente; i++ {
        resultado *= base
    }
    return resultado
}

func main() {
    var numero int
    fmt.Print("Digite um número inteiro: ")
    fmt.Scanln(&numero)

    if numero%2 == 0 {
        fmt.Println("O número", numero, "é par.")
    } else {
        fmt.Println("O número", numero, "é ímpar.")
    }
}