package main

import "fmt"

func main(){
	verificaPariedade()
}

func verificaPariedade(){
	var n1 int
	fmt.Println("Digite o numero a ser verificado: ")
	fmt.Scanln(&n1)

	if n1 % 2 == 0{
		fmt.Println("O numero e par")
	} else{
		fmt.Println("O numero e impar")
	}
}