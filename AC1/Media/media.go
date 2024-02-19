package main

import "fmt"

func main() {
    var n1, n2, n3 float64

    fmt.Print("Digite os três números: ")
    fmt.Scanln(&n1, &n2, &n3)

    media := calculaMedia(n1, n2, n3)
    fmt.Printf("A média é: %.2f\n", media)
}

func calculaMedia(valores ...float64) float64 {
    total := 0.0

    for _, valor := range valores {
        total += valor
    }

    if len(valores) > 0 {
        return total / float64(len(valores))
    } else {
        return 0.0
    }
}
