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
		showHands(initialLoop, allPlayers, deck)
		//players turns
		allPlayers, deck = playerTurn(allPlayers, deck)

		//revealing dealers second card and total
		dealer.dealer = false
		fmt.Printf("Dealer: %v", printTotal(*dealer))

		//dealers turn
		dealer, deck = dealerTurn(dealer, deck)
		dealer.checkBust()

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

func showHands(initialLoop bool, allPlayers []*Hand, deck Deck) {
	dealer := allPlayers[len(allPlayers)-1]
	var newHands []*Hand
	numPlayers := len(allPlayers) - 1

	if initialLoop {
		initialLoop = false
		fmt.Println("************************")
		for i, v := range allPlayers {
			if i == len(allPlayers)-1 {
				continue
			}
			fmt.Printf("Player %v: %v\n", i+1, printTotal(*v))
		}
		fmt.Println("Dealer:", dealer.String())
		fmt.Println("************************")
	} else {
		//subsequent rounds require the players hands to be replaced with newly dealt hands
		newHands, deck = deal(deck, numPlayers)
		fmt.Println("************************")
		for i, v := range allPlayers {
			v.hand = newHands[i].hand
			if i == len(allPlayers)-1 {
				v.dealer = true
				continue
			}
			fmt.Printf("Player %v: %v\n", i+1, printTotal(*v))
		}
		fmt.Println("Dealer:", dealer.String())
		fmt.Println("************************")
	}
}
