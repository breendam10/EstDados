package main

import "fmt"


func MoverTorre(n int, origem, destino, auxiliar string) {
	if n == 1 {
		fmt.Printf("Move 1 de %s para %s\n", origem, destino)
		return
	}

	MoverTorre(n-1, origem, auxiliar, destino)
	fmt.Printf("Move %d de %s para %s\n", n, origem, destino)
	MoverTorre(n-1, auxiliar, destino, origem)
}


func achapar(arr []int, target int) (int, int) {

	seen := make(map[int]int)

	for i, num := range arr {

		complemento := target - num


		if j, ok := seen[complemento]; ok {

			return j, i
		}


		seen[num] = i
	}

	return -1, -1
}


func main() {

	numDiscos := 3
	origem, destino, auxiliar := "A", "C", "B"

	MoverTorre(numDiscos, origem, destino, auxiliar)


	arr := []int{1, 3, 5, 7, 9, 11, 14, 15, 20}
	target := 10

	i, j := achapar(arr, target)
	if i == -1 && j == -1 {
		fmt.Println("Tem par.")
	} else {
		fmt.Printf("Não tem par: %d, %d\n", arr[i], arr[j])
	}
}







