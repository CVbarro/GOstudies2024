package main

import "fmt"

func main() {
    var competidores, folhasCompradas, folhasPorCompetidor int
    fmt.Println("Digite o número de competidores, a quantidade de folhas compradas e a quantidade de folhas por competidor:")
    fmt.Scanln(&competidores, &folhasCompradas, &folhasPorCompetidor)

    totalFolhasNecessarias := competidores * folhasPorCompetidor

    if totalFolhasNecessarias <= folhasCompradas {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}
