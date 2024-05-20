package main

import "fmt"

func determinarCategoria(K int)int {
	categorias := []int{1, 3,5, 10, 25,50, 100}
	for _, categoria := range categorias {
		if K<= categoria {
			return categoria
		}
	}
	return 100 
}
func main() {
	var K int
	fmt.Println("Digite a colocação:")
	fmt.Scanln(&K)
	categoria := determinarCategoria(K)
	fmt.Printf("Top %d\n", categoria)
}