package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Println("Digite o número de viagens de Dona Parcinova ao mercado:")
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var M int
		fmt.Println("\nDigite a quantidade de produtos disponíveis no mercado:")
		fmt.Scanln(&M)

		produtos := make(map[string]float64)
		for j := 0; j < M; j++ {
			var nomeProduto string
			var preco float64
			fmt.Println("Digite o nome do produto e o preço por unidade ou Kg:")
			fmt.Scanln(&nomeProduto, &preco)
			produtos[nomeProduto] = preco
		}

		var P int
		fmt.Println("\nDigite a quantidade de produtos que Dona Parcinova deseja comprar:")
		fmt.Scanln(&P)

		total := 0.0
		for j := 0; j < P; j++ {
			var nomeItem string
			var quantidade int
			fmt.Println("Digite o nome do produto e a quantidade desejada:")
			fmt.Scanln(&nomeItem, &quantidade)
			total += float64(quantidade) * produtos[nomeItem]
		}

		fmt.Printf("\nR$ %.2f\n", total)
	}
}
