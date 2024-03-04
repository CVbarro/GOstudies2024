package main

import (
	"fmt"
	"math"
	"geometria"
)

func main(){
	meuVetor := criarVetor()
	iniciarVetor(meuVetor)
	mostrarVetor(meuVetor)

	minhaPalavra :=pegarPalavra()
	palavraInvertida := inverterPalavra(minhaPalavra)
	println(palavraInvertida)

	ponto := Ponto{x: 3, y: 4}
	distancia := ponto.DistanciaOrigem()
	fmt.Println("A distância do ponto (%2f, %2f) até a origem é %2f\n", ponto.x, ponto.y, distancia)

	var largura, altura float64
	fmt.Print("Digite a largura do retângulo: ")
	fmt.Scan(&largura)
	fmt.Print("Digite a altura do retângulo: ")
	fmt.Scan(&altura)

	area := geometria.CalcularArea(largura, altura)
	perimetro := geometria.CalcularPerimetro(largura, altura)

	
	fmt.Printf("Área do retângulo: %.2f\n", area)
	fmt.Printf("Perímetro do retângulo: %.2f\n", perimetro)
}

func criarVetor() []int{
	vetor := make([]int, 10)
	return vetor
}

func iniciarVetor(vetor []int){
	for i := 1; i < 10; i++{
		vetor[i] = (i + 1) * 2
	}
}

func mostrarVetor(vetor []int){
	for _, elemento := range vetor {
		fmt.Println(elemento)
	}

}

func pegarPalavra() string {
	var palavra string

	fmt.Print("Digite a palavra a ser invertida: ")
	fmt.Scan(&palavra)
	return palavra
}

func inverterPalavra(palavra string) string {
	letras := []rune(palavra)
	comprimentoPalavra := len(letras)

	for i, j := 0, comprimentoPalavra - 1; i < j; i, j = i + 1, j - 1 {
		letras[i], letras[j] = letras[j], letras[i]
	}

	return string(letras)
}

type Ponto struct{
	x , y float64
}

func (p Ponto) DistanciaOrigem() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}