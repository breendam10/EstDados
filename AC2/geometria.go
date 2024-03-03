package main

import "fmt"


func Area(base, altura float64) float64 {
    return base * altura
}


func Perimetro(base, altura float64) float64 {
    return 2 * (base + altura)
}


func CalcularAreaPerimetro(base, altura float64) {
    area := Area(base, altura)
    perimetro := Perimetro(base, altura)

    fmt.Printf("Área do retângulo: %.2f\n", area)
    fmt.Printf("Perímetro do retângulo: %.2f\n", perimetro)
}
