package main

import "fmt"

func minutosParaDormir(horaAtual, minutoAtual, horaAlarme, minutoAlarme int) int {
	horaAtualEmMinutos := horaAtual*60 + minutoAtual
	horaAlarmeEmMinutos := horaAlarme*60 + minutoAlarme
	if horaAlarmeEmMinutos <= horaAtualEmMinutos {
		horaAlarmeEmMinutos += 24 * 60
	}
	return horaAlarmeEmMinutos - horaAtualEmMinutos
}

func main() {
	var horaAtual, minutoAtual, horaAlarme, minutoAlarme int
	for {
		fmt.Scan(&horaAtual, &minutoAtual, &horaAlarme, &minutoAlarme)
		if horaAtual == 0 && minutoAtual == 0 && horaAlarme == 0 && minutoAlarme == 0 {
			break 
		}
		minutos := minutosParaDormir(horaAtual, minutoAtual, horaAlarme, minutoAlarme)
		fmt.Println(minutos)
	}
}
