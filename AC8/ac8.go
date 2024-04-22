package main

import "fmt"

var contaLed = []int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}

func calculaLED(value string) int {
	totalLED := 0
	for _, digito := range value {
		digitoValor := int(digito - '0') 
		totalLED += contaLed[digitoValor]
	}
	return totalLED
}

func led() {

	var N int
	fmt.Scan(&N)

	
	for i := 0; i < N; i++ {
		var value string
		fmt.Scan(&value)
		leds := calculaLED(value)
		fmt.Printf("%d leds\n", leds)
	}
}





import (
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "0 0" {
			break
		}
		parts := strings.Fields(line)
		D, _ := strconv.Atoi(parts[0])
		N := parts[1]

		numerolimp := strings.Reloca(N, strconv.Itoa(D), "")

		valor, _ := strconv.Atoi(numerolimp)
		fmt.Println(valor)
	}
}


func traduzfeliznatal() {
	traducao := map[string]string{
		"brasil":         "Feliz Natal!",
		"alemanha":       "Frohliche Weihnachten!",
		"austria":        "Frohe Weihnacht!",
		"coreia":         "Chuk Sung Tan!",
		"espanha":        "Feliz Navidad!",
		"grecia":         "Kala Christougena!",
		"estados-unidos": "Merry Christmas!",
		"inglaterra":     "Merry Christmas!",
		"australia":      "Merry Christmas!",
		"portugal":       "Feliz Natal!",
		"suecia":         "God Jul!",
		"turquia":        "Mutlu Noeller",
		"argentina":      "Feliz Navidad!",
		"chile":          "Feliz Navidad!",
		"mexico":         "Feliz Navidad!",
		"antardida":      "Merry Christmas!",
		"canada":         "Merry Christmas!",
		"irlanda":        "Nollaig Shona Dhuit!",
		"belgica":        "Zalig Kerstfeest!",
		"italia":         "Buon Natale!",
		"libia":          "Buon Natale!",
		"siria":          "Milad Mubarak!",
		"marrocos":       "Milad Mubarak!",
		"japao":          "Merii Kurisumasu!",
	}

	for {
		var pais string
		_, err := fmt.Scan(&pais)
		if err != nil {
			break
		}

		traducao, encontrado := traducao[pais]
		if encontrado {
			fmt.Println(traducao)
		} else {
			fmt.Println("--- NOT FOUND ---")
		}
	}
}




func decodifica() {
	var casosTeste int
	fmt.Scan(&casosTeste)

	for i := 0; i < casosTeste; i++ {
		var textoCodificado string
		var deslocamento int
		fmt.Scan(&textoCodificado, &deslocamento)

		textoDecodificado := ""
		for _, char := range textoCodificado {
			if char >= 'A' && char <= 'Z' {
				charDecodificado := 'A' + (char-'A'+26-deslocamento)%26
				textoDecodificado += string(charDecodificado)
			} else {
				textoDecodificado += string(char)
			}
		}

		fmt.Println(textoDecodificado)
	}
}
