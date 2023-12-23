package main

import "fmt"

func main() {

	//Inicialização de variavel
	cards := newDeck()

	//Atribuindo valores do retorno da função para variaveis
	hand, remainingCards := deal(cards, 5)

	fmt.Println("Player Hand")
	hand.print()
	fmt.Println("Remaining Cards")
	remainingCards.print()

}
