package main

import (
	"fmt"
	"strings"

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

	var move string
	var drawnCard Card

	//should look into making this into game loop function?
	for i := 0; i < len(allPlayers)-1; i++ {
		iter := 0
		fmt.Printf("Player %v's turn\n", i+1)
		for {
			if iter > 0 {
				fmt.Printf("Player %v: %v\n", i+1, printTotal(*allPlayers[i]))
			}
			if checkBlackjack(*allPlayers[i]) {
				fmt.Println("Blackjack!")
				break
			}
			if checkBust(*allPlayers[i]) {
				fmt.Println("Bust. Dealer wins.")
				break
			}
			fmt.Println("Hit or stand?")
			iter++
			fmt.Scanf("%s", &move)
			move = strings.ToLower(move)

			if move == "hit" {
				drawnCard, deck = drawCard(deck)
				fmt.Println(drawnCard.String())
				fmt.Println()
				allPlayers[i].hand = append(allPlayers[i].hand, drawnCard)
				continue
			}
			fmt.Printf("Player %v: %v\n", i+1, printTotal(*allPlayers[i]))

			if move == "stand" {
				break
			}

			fmt.Println("Invalid entry")
		}
	}
	//display dealer hand

	dealer.dealer = false
	fmt.Printf("Dealer : %v\n", printTotal(*dealer))

	//dealer turn
	dealer = dealerTurn(dealer, deck)
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
