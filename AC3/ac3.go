
package main

import (
	"fmt"
	"os"
)

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func adicionarContato(contato Contato, arrayContatos []Contato) []Contato {
	for i := range arrayContatos {
		if arrayContatos[i].Nome == "" && arrayContatos[i].Email == "" {
			arrayContatos[i] = contato
			break
		}
	}
	return arrayContatos
}

func excluirContato(arrayContatos []Contato) []Contato {
	for i := len(arrayContatos) - 1; i >= 0; i-- {
		if arrayContatos[i].Nome != "" || arrayContatos[i].Email != "" {
			arrayContatos[i] = Contato{}
			break
		}
	}
	return arrayContatos
}

func main() {
	contatos := make([]Contato, 5)
	for {
		fmt.Println("\n1. Adicionar Contato")
		fmt.Println("2. Excluir Último Contato")
		fmt.Println("3. Sair")

		var opcao int
		fmt.Print("Escolha uma opção (1-3): ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			var nome, email string
			fmt.Print("Digite o nome do contato: ")
			fmt.Scan(&nome)
			fmt.Print("Digite o e-mail do contato: ")
			fmt.Scan(&email)

			novoContato := Contato{Nome: nome, Email: email}
			contatos = adicionarContato(novoContato, contatos)
			fmt.Println("Contato adicionado com sucesso!")

		case 2:
			contatos = excluirContato(contatos)
			fmt.Println("Último contato excluído com sucesso!")

		case 3:
			fmt.Println("Saindo...")
			os.Exit(0)

		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

