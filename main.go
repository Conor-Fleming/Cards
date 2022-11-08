package main

import (
	"fmt"
	"strings"

	"github.com/Conor-Fleming/deck/deck"
)

type Card = deck.Card

type Hand []Card

type Deck []Card

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

	// get input for how many players
	// ?

	// call deal function with that amount
	allPlayers, deck := deal(deck, 5)
	for i, v := range allPlayers {
		if i == len(allPlayers)-1 {
			continue
		}
		fmt.Printf("Player %v: %v\n", i+1, v.String())
	}
	fmt.Println("Dealer:", allPlayers[len(allPlayers)-1].String())
}

func deal(deck Deck, players int) ([]*Hand, Deck) {
	// creating players and dealer hands
	HandArr := make([]*Hand, 0)

	for i := 0; i < players; i++ {
		var player Hand
		HandArr = append(HandArr, &player)
	}
	var dealer Hand
	HandArr = append(HandArr, &dealer)

	// iterate through players and dealer twice and deal a card on each iteration
	// similar to how a real game would be dealt
	for i := 0; i < 2; i++ {
		for _, v := range HandArr {
			card := deck[0]
			deck = deck[1:]
			*v = append(*v, card)
		}
	}

	return HandArr, deck
}

func drawCard(deck Deck) (Card, Deck) {
	card := deck[0]
	deck = deck[1:]
	return card, deck
}
