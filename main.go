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

	//get input for how many players
	// ?

	//call deal function with that amount
	allPlayers := deal(deck, 1)

	for i, v := range allPlayers {
		if i == len(allPlayers)-1 {
			continue
		}
		fmt.Printf("Player %v: %v\n", i+1, v.String())
	}
	fmt.Println("Dealer: ", allPlayers[len(allPlayers)-1].String())
}

func deal(deck []deck.Card, players int) []Hand {
	//creating players and dealer hands
	HandArr := make([]Hand, players+1)
	for i := 0; i < players; i++ {
		var player Hand
		HandArr = append(HandArr, player)
	}
	var dealer Hand
	HandArr = append(HandArr, dealer)

	for _, v := range HandArr {
		card := deck[0]
		deck = deck[1:]
		v = append(v, card)
	}

	return HandArr
}

func drawCard(deck []deck.Card) (Card, Hand) {
	card := deck[0]
	deck = deck[1:]
	return card, deck
}
