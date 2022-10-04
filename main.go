package main

import (
	"github.com/Conor-Fleming/deck/deck"
)

type Hand struct {
}

type Player struct {
}

// version 1
func main() {
	//Decks sets the amount of extra decks to add to the intial deck
	decks := 2

	//deal cards
	deck := deck.New(deck.ExtraDecks(decks), deck.Shuffle)

	//player can hit or stand

}
