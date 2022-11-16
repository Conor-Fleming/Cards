package main

import (
	"fmt"
	"strings"
)

func playerTurn(allPlayers []*Hand, deck Deck) ([]*Hand, Deck) {
	var move string
	var card Card

	for i := 0; i < len(allPlayers)-1; i++ {
		iter := 0
		fmt.Printf("Player %v's turn\n", i+1)
		for {
			if iter > 0 {
				fmt.Printf("Player %v: %v\n", i+1, printTotal(*allPlayers[i]))
			}
			if allPlayers[i].checkBlackjack() {
				fmt.Println("Blackjack!")
				break
			}
			if allPlayers[i].checkBust() {
				fmt.Println("Bust.")
				allPlayers[i].bust = true
				break
			}
			fmt.Println("Hit or stand?")
			iter++
			fmt.Scanf("%s", &move)
			move = strings.ToLower(move)

			if move == "hit" {
				card, deck = drawCard(deck)
				fmt.Println("Player hits...", card.String())
				allPlayers[i].hand = append(allPlayers[i].hand, card)
				continue
			}
			//fmt.Printf("Player %v: %v\n", i+1, printTotal(*allPlayers[i]))

			if move == "stand" {
				break
			}

			fmt.Println("Invalid entry")
		}
	}
	return allPlayers, deck
}
