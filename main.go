package main

import (
	"fmt"

	"github.com/Conor-Fleming/deck/deck"
)

type Card = deck.Card

type Player struct {
	Dealer bool
	Hand   []deck.Card
}

// version 1
func main() {
	//Decks sets the amount of extra decks to add to the intial deck
	decks := 2

	//deal cards
	deck := deck.New(deck.ExtraDecks(decks), deck.Shuffle)
	dealer, player1 := deal(deck)
	fmt.Println("The dealers hand:\n", dealer[1:])
	fmt.Println("Player 1's hand:\n", player1)

	//player can hit or stand

}

func deal(deck []Card) ([]Card, []Card) {
	var p1 []Card
	var dealer []Card
	p1 = append(p1, deck[0])
	dealer = append(dealer, deck[1])
	p1 = append(p1, deck[2])
	dealer = append(dealer, deck[3])
	return dealer, p1
}
