package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calcularHipotenusa(N, H, C int) float64 {
	alturaTotal := float64(N * H)
	profundidadeTotal := float64(N * C)
	return math.Sqrt(alturaTotal*alturaTotal + profundidadeTotal*profundidadeTotal)
}

func calcularArea(N, H, C, L int) float64 {
	hipotenusa := calcularHipotenusa(N, H, C)
	areaCm := hipotenusa * float64(L)
	return areaCm * 0.0001
}

func processarEntrada() []float64 {
	var resultados []float64
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		N, _ := strconv.Atoi(scanner.Text())
		if !scanner.Scan() {
			break
		}
		dimensoes := strings.Fields(scanner.Text())
		H, _ := strconv.Atoi(dimensoes[0])
		C, _ := strconv.Atoi(dimensoes[1])
		L, _ := strconv.Atoi(dimensoes[2])

		resultado := calcularArea(N, H, C, L)
		resultados = append(resultados, resultado)
	}
	return resultados
}

func main() {
	resultados := processarEntrada()
	for _, resultado := range resultados {
		fmt.Printf("%.4f\n", resultado)
	}
}
