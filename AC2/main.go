package main

import (
    "fmt"
)

func main() {
    var base, altura float64
    fmt.Print("Digite a base do retângulo: ")
    fmt.Scanln(&base)
    fmt.Print("Digite a altura do retângulo: ")
    fmt.Scanln(&altura)

    CalcularAreaPerimetro(base, altura)
}
