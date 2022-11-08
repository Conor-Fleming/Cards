package main

import (
	"fmt"
	"strings"

	"github.com/Conor-Fleming/deck/deck"
)

type Card = deck.Card

type Hand []Card

func (h Hand) String() string {
	stringsArr := make([]string, len(h))
	for i := range h {
		stringsArr[i] = h[i].String()
	}
	return strings.Join(stringsArr, ", ")
}

// version 1
func main() {
	//Decks sets the amount of extra decks to add to the intial deck
	deck := deck.New(deck.ExtraDecks(2), deck.Shuffle)

	player, deck := deal(deck)
	fmt.Println("Player hand: ", player.String())
	dealer, deck := deal(deck)
	fmt.Println("Dealer Hand: ", dealer.String())

}

func deal(deck []deck.Card) (Hand, []deck.Card) {
	var hand Hand
	for i := 0; i < 2; i++ {
		hand = append(hand, deck[i])
		deck = deck[1:]
	}
	return hand, deck
}

func drawCard(deck []deck.Card) (Card, []deck.Card) {
	card := deck[0]
	deck = deck[1:]
	return card, deck
}
