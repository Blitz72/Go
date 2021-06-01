package main

import "fmt"

func main() {
	// card := "Ace of Spades"
	cards := deck{newCard(), "Ace of Diamonds"}
	card := newCard()
	cards = append(cards, "Six of Spades")
	fmt.Println(card)
	fmt.Println(cards)
	fmt.Println(len(cards))

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

// var decksize int

// func main() {
// 	var card string
// 	// var card string = "Ace of Spades"
// 	// card := "Ace of Spades"
// 	card = "5 of Diamonds"
// 	decksize = 52
// 	fmt.Println(card)
// 	fmt.Println(decksize)
// }
