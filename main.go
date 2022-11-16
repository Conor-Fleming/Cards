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
	// ?

	// call deal function with that amount
	allPlayers, deck := deal(deck, 1)

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

	dealer.dealer = false
	fmt.Printf("Dealer: %v", printTotal(*dealer))

	//dealer turn
	dealer, deck = dealerTurn(dealer, deck)
	if checkBust(*dealer) {
		fmt.Println("Dealer busts.")
	} else {
		//compare scores //create function for this that contains checking logic
		switch findWinner(*dealer, *allPlayers[0]) {
		case "dealer":
			fmt.Println("Dealer wins!")
		case "player":
			fmt.Println("Player wins!")
		case "push":
			fmt.Println("Push. It was a draw.")
		}
	}
	//run through players with bust == false to compare scores
}
