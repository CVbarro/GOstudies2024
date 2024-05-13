package main

import (
	"fmt"
)

func main() {
	var distancia int
	fmt.Println("Digite a distância em metros:")
	fmt.Scanln(&distancia)

	tempo := distancia * 2

	fmt.Println("Tempo necessário para percorrer a distância:", tempo, "minutos")
}
