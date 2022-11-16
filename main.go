package main

import (
	"fmt"

	"github.com/Conor-Fleming/deck/deck"
)

// version 1
func main() {
	//Decks sets the amount of extra decks to add to the intial deck
	deck := deck.New(deck.ExtraDecks(2), deck.Shuffle)

	// get input for how many players
	fmt.Println("Welcome to Blackjack")

	fmt.Println("How many players?")
	var numPlayers int
	fmt.Scanf("%v", &numPlayers)

	// call deal function with that amount
	allPlayers, deck := deal(deck, numPlayers)

	dealer := allPlayers[len(allPlayers)-1]

	for i, v := range allPlayers {
		if i == len(allPlayers)-1 {
			continue
		}
		fmt.Printf("Player %v: %v\n", i+1, printTotal(*v))
	}
	fmt.Println("Dealer:", dealer.String())
	fmt.Println()

	//should look into making this into game loop function?
	allPlayers, deck = playerTurn(allPlayers, deck)
	//display dealer hand
	fmt.Println(len(allPlayers) - 1)

	dealer.dealer = false
	fmt.Printf("Dealer: %v", printTotal(*dealer))

	//dealer turn
	dealer, deck = dealerTurn(dealer, deck)
	if dealer.checkBust() {
		fmt.Println("Dealer busts.")
	}

	findWinner(dealer, allPlayers)
}
