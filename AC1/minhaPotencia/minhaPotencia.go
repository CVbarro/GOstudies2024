package main

import "fmt"

func main() {
    base := 2
    expoente := 3
    resultado := minhaPotencia(base, expoente)

    fmt.Printf("%d elevado a %d é igual a %d\n", base, expoente, resultado)
}

func minhaPotencia(base, expoente int) int {
    resultado := 1

    for i := 0; i < expoente; i++ {
        resultado *= base
    }

    return resultado
}
