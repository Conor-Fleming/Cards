package main

import (
	"fmt"
	"time"
)

func dealerTurn(dealer *Hand, deck Deck) (*Hand, Deck) {
	//check score, on 16 or less hit on soft 17 hit
	var card Card
	for {
		total, soft := getTotal(*dealer)
		if (total == 17 && !soft) || total > 16 {
			break
		}
		card, deck = drawCard(deck)
		dealer.hand = append(dealer.hand, card)
		total, _ = getTotal(*dealer)
		fmt.Printf("Dealer hits... %v Total: %v\n", card.String(), total)
		time.Sleep(time.Second * 2)
	}
	return dealer, deck
}
