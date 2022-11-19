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
	fmt.Println()

	deck := deck.New(deck.ExtraDecks(2), deck.Shuffle)
	allPlayers, deck := deal(deck, numPlayers)
	dealer := allPlayers[len(allPlayers)-1]

	//create game loop, ask to play again after each round
	cont := true
	initialLoop := true

	for cont {
		var newHands []*Hand
		if initialLoop {
			initialLoop = false
			for i, v := range allPlayers {
				if i == len(allPlayers)-1 {
					continue
				}
				fmt.Printf("Player %v: %v\n", i+1, printTotal(*v))
			}
			fmt.Println("Dealer:", dealer.String())
		} else {
			//subsequent rounds require the players hands to be replaced with newly dealt hands
			newHands, deck = deal(deck, numPlayers)
			for i, v := range allPlayers {
				v.hand = newHands[i].hand
				if i == len(allPlayers)-1 {
					v.dealer = true
					continue
				}
				fmt.Printf("Player %v: %v\n", i+1, printTotal(*v))
			}
			fmt.Println("Dealer:", dealer.String())
		}
		//players turns
		allPlayers, deck = playerTurn(allPlayers, deck)

		//revealing dealers second card and total
		dealer.dealer = false
		fmt.Printf("Dealer: %v", printTotal(*dealer))

		//dealers turn
		dealer, deck = dealerTurn(dealer, deck)
		if dealer.checkBust() {
			dealer.bust = true
			fmt.Println("Dealer busts.")
		}

		//compare scores and track wins and losses
		findWinner(dealer, allPlayers)

		//ask user if they want to keep playing
		//if yes loop continues, if no stats will display for session
		fmt.Println("Play again? Type yes or no")
		var input string
		fmt.Scanf("%v", &input)
		if input == "no" {
			cont = false
			fmt.Println()
			fmt.Println(stats(allPlayers))
		}
	}
}
