package q4

import "fmt"

//Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
//deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
//lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.
//
//Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
//ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
//estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
//Caso a lista possua apenas um elemento, a função deve retornar 3.

func ClassifyPrices(prices []int) (int, error) {
	if len(prices) == 0 {
		return 0, fmt.Errorf("Lista vazia")
	}
	if len(prices) == 1 {
		return 3, nil
	}
	ordem := 3
	crescente := true
	for i := 1; i < len(prices); i++ {
		atual1 := prices[i]
		anterior1 := prices[i-1]
		if atual1 < anterior1 {
			crescente = false
			break
		}
		anterior1 = atual1
	}
	decrescente := true
	for y := 1; y < len(prices); y++ {
		atual2 := prices[y]
		anterior2 := prices[y-1]
		if atual2 > anterior2 {
			decrescente = false
			break
		}
		anterior2 = atual2
	}
	if crescente {
		ordem = 1
	}
	if decrescente {
		ordem = 2
	}
	return ordem, nil
}
