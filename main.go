package main

// Desenvolvido por: Gabriel Feliciano da Silva para o Teste Transdata Desenvolvedor Pleno

import (
	"fmt"
	"sort"
)

// calculateNotes calcula a quantidade mínima de notas para o valor monetário fornecido.
// A função itera sobre as notas disponíveis (de maior para menor) e determina quantas de cada são necessárias.
func calculateNotes(amount int) (map[int]int, bool) {
	notes := []int{200, 100, 50, 20, 10, 5}
	result := make(map[int]int)

	// Se o valor não for múltiplo da menor nota, é impossível sacar.
	if amount%5 != 0 {
		return nil, false
	}

	remainingAmount := amount

	for _, note := range notes {
		if remainingAmount >= note {
			count := remainingAmount / note
			result[note] = count
			remainingAmount %= note
		}
	}

	// Se sobrou algum valor, algo deu errado.
	if remainingAmount != 0 {
		return nil, false
	}

	return result, true
}

func main() {
	var monetaryValue int

	fmt.Print("Digite o valor monetário para o saque: ")
	_, err := fmt.Scan(&monetaryValue)
	if err != nil {
		fmt.Println("Erro ao ler o valor. Por favor, insira um número inteiro.")
		return
	}

	notesDistribution, possible := calculateNotes(monetaryValue)

	if possible {
		// Se o valor for válido, exibe a quantidade de cada nota.
		// Exemplo: Entrada 770 => saída (3 notas de 200, 1 nota de 100, 1 nota de 50 e 1 nota de 20).
		fmt.Printf("Para sacar R$%d, você receberá:\n", monetaryValue)

		// Ordena as chaves (notas) para exibição consistente
		var sortedNotes []int
		for note := range notesDistribution {
			sortedNotes = append(sortedNotes, note)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(sortedNotes)))

		for _, note := range sortedNotes {
			count := notesDistribution[note]
			fmt.Printf("%d nota(s) de R$%d\n", count, note)
		}
	} else {
		// Mensagem exibida caso o valor esteja indisponível para saque.
		fmt.Printf("O valor de R$%d está indisponível para saque com as notas existentes (200, 100, 50, 20, 10, 5).\n", monetaryValue)
	}
}
