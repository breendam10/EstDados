package main

import "fmt"

// Função para determinar a menor categoria
func determinarCategoria(K int) int {
	categorias := []int{1, 3, 5, 10, 25, 50, 100}
	for _, categoria := range categorias {
		if K <= categoria {
			return categoria
		}
	}
	return 100 // Se não se encaixar em nenhuma categoria, retorna 100
}

func main() {
	var K int
	fmt.Println("Digite o número da colocação:")
	fmt.Scanln(&K)

	// Determinando a categoria e imprimindo a saída
	categoria := determinarCategoria(K)
	fmt.Printf("Top %d\n", categoria)
}
