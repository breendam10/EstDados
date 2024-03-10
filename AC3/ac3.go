package main


import "fmt"

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func AdicionarContato(contato Contato, arrayContatos []*Contato) {
	for i, c := range arrayContatos {
		if c == nil {
			arrayContatos[i] = &contato
			return
		}
	}
	fmt.Println(" Não é possível adicionar mais contatos.")
}

func ExcluirContato(arrayContatos []*Contato) {
	for i := len(arrayContatos) - 1; i >= 0; i-- {
		if arrayContatos[i] != nil {
			arrayContatos[i] = nil
			return
		}
	}
	fmt.Println(" Não há contatos para excluir.")
}

func main() {
	arrayContatos := make([]*Contato, 5)

	for {
		fmt.Println("\nEscolha uma opção:")
		fmt.Println("1 - Adicionar um novo contato")
		fmt.Println("2 - Excluir o último contato")

		var opcao int
		fmt.Print("Opção: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			var nome, email string
			fmt.Println("\nDigite o nome:")
			fmt.Scanln(&nome)

			fmt.Println("Digite o email:")
			fmt.Scanln(&email)

			contato := Contato{Nome: nome, Email: email}
			AdicionarContato(contato, arrayContatos)
			fmt.Println("Operação concluida.")
		case 2:
			ExcluirContato(arrayContatos)
			fmt.Println("Operação concluida.")
		default:
			fmt.Println("Opção inválida.")
		}

		fmt.Println("\nLista de contatos:")
		for i, c := range arrayContatos {
			if c != nil {
				fmt.Printf("%d: %s - %s\n", i+1, c.Nome, c.Email)
			} else {
				fmt.Printf("%d: Vazio\n", i+1)
			}
		}
	}
}
