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
		total, _ := getTotal(*allPlayers[i])
		fmt.Printf("\nPlayer %v's turn (Total: %v)\n", i+1, total)
		for {
			if iter > 0 {
				fmt.Printf("Player %v: %v", i+1, printTotal(*allPlayers[i]))
			}
			iter++

			if scoreCheck(allPlayers[i]) {
				break
			}

			fmt.Println("Hit or stand?")
			fmt.Scanf("%s", &move)
			move = strings.ToLower(move)
			if move == "hit" {
				card, deck = drawCard(deck)
				printHit(card, "Player")
				allPlayers[i].hand = append(allPlayers[i].hand, card)
				continue
			}
			if move == "stand" {
				break
			}
			fmt.Println("Invalid entry")
		}
	}
	return allPlayers, deck
}

func scoreCheck(h *Hand) bool {
	if h.checkBust() {
		h.bust = true
		return true
	}
	if h.checkBlackjack() {
		return true
	}
	return false
}
