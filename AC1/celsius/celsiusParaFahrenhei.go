package main

import "fmt"

func main() {
    temperaturaCelsius := 25.0
    temperaturaFahrenheit := converteCelsiusParaFahrenheit(temperaturaCelsius)

    fmt.Printf("%.2f graus Celsius é equivalente a %.2f graus Fahrenheit.\n", temperaturaCelsius, temperaturaFahrenheit)
}

func converteCelsiusParaFahrenheit(celsius float64) float64 {
    fahrenheit := (celsius * 9/5) + 32
    return fahrenheit
}