package main

import (
	"fmt"
)


func hanoi(n int, primeiraTorre, torreMeio, torreObjetivo string) {
	if n > 0 {
		hanoi(n-1, primeiraTorre, torreObjetivo, torreMeio)
		fmt.Printf("Mova o disco %d de %s para %s\n", n, primeiraTorre, torreObjetivo)
		hanoi(n-1, torreMeio, primeiraTorre, torreObjetivo)
	}
}

func somaAlvo(array []int, alvo int) (int, int) {
	i, j := 0, len(array)-1

	for i < j {
		soma := array[i] + array[j]
		if soma == alvo {
			return array[i], array[j]
		} else if soma < alvo {
			i++
		} else {
			j--
		}
	}

	return -1, -1
}

func main() {
	fmt.Println("Solução para a Torre de Hanói:")
	hanoi(3, "A", "B", "C")

	fmt.Println("\nEncontrar par com soma alvo:")
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	alvo := 12
	resultado1, resultado2 := somaAlvo(array, alvo)

	if resultado1 != -1 && resultado2 != -1 {
		fmt.Printf("Par encontrado: (%d, %d)\n", resultado1, resultado2)
	} else {
		fmt.Println("Nenhum par encontrado.")
	}
}
