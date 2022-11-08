package main

import (
	"fmt"

	"github.com/Conor-Fleming/deck/deck"
)

type Card = deck.Card

// version 1
func main() {
	//Decks sets the amount of extra decks to add to the intial deck
	deck := deck.New(deck.ExtraDecks(2), deck.Shuffle)
	test, deck := drawCard(deck)
	fmt.Println(test)

}

func drawCard(deck []deck.Card) (Card, []deck.Card) {
	card := deck[0]
	deck = deck[1:]
	return card, deck
}
