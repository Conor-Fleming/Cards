package main

import (
	"fmt"

	"github.com/Conor-Fleming/deck/deck"
)

// version 1
func main() {
	// get input for how many players
	fmt.Println("Welcome to Blackjack")
	fmt.Println("How many players?")
	var numPlayers int

	fmt.Scanf("%v", &numPlayers)

	deck := deck.New(deck.ExtraDecks(2), deck.Shuffle)
	allPlayers, deck := deal(deck, numPlayers)
	dealer := allPlayers[len(allPlayers)-1]

	//create game loop, ask to play again after each round
	cont := true
	initial := true
	for cont {
		var test []*Hand
		if !initial {
			test, deck = deal(deck, numPlayers)
		}
		initial = false
		// call deal function with that amount

		for i, v := range allPlayers {
			v.hand = test[i].hand
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
			dealer.bust = true
			fmt.Println("Dealer busts.")
		}

		findWinner(dealer, allPlayers)

		fmt.Println("Play again? Type yes or no")
		var input string
		fmt.Scanf("%v", &input)
		if input == "no" {
			cont = false
			fmt.Println()
			//win loss data is not persiting
			fmt.Println(stats(allPlayers))
		}
	}
}
