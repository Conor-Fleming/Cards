package main

import (
	"fmt"
	"strings"

	"github.com/Conor-Fleming/deck/deck"
)

type Card = deck.Card

type Hand struct {
	hand   []Card
	dealer bool
	bust   bool
}

type Deck []Card

func (h Hand) String() string {
	stringsArr := make([]string, len(h.hand))
	if !h.dealer {
		for i := range h.hand {
			stringsArr[i] = h.hand[i].String()
		}
		return strings.Join(stringsArr, ", ")
	}
	return h.hand[1].String()
}

func checkBlackjack(hand Hand) bool {
	total, _ := getTotal(hand)
	if total == 21 {
		return true
	}
	return false
}

func checkBust(hand Hand) bool {
	total, _ := getTotal(hand)
	if total > 21 {
		hand.bust = true
		return true
	}
	return false
}

func printTotal(value Hand) string {
	total, soft := getTotal(value)
	if soft {
		return fmt.Sprintf("%v  Total: %v / %v\n", value.String(), total-10, total)
	}

	return fmt.Sprintf("%v  Total: %v\n", value.String(), total)
}

func findWinner(dealer Hand, player Hand) bool {
	dealerScore, _ := getTotal(dealer)
	playerScore, _ := getTotal(player)
	if dealerScore > playerScore {
		return true
	}
	return false
}

func getTotal(hand Hand) (int, bool) {
	var total int
	var ace bool
	var soft bool
	for _, v := range hand.hand {
		if v.Value == 1 { //Ace card
			ace = true
		}
		if v.Value > 10 {
			v.Value = 10
		}
		//using value of card (cards over 10 contain values 11, 12, 13) messing up math for blackjack
		total += int(v.Value)
	}

	if ace && total < 12 {
		total += 10
		soft = true
	}

	return total, soft
}

func deal(deck Deck, players int) ([]*Hand, Deck) {
	// creating players and dealer hands
	HandArr := make([]*Hand, 0)

	for i := 0; i < players; i++ {
		player := Hand{
			hand:   nil,
			dealer: false,
		}
		HandArr = append(HandArr, &player)
	}
	dealer := Hand{
		hand:   nil,
		dealer: true,
	}
	HandArr = append(HandArr, &dealer)

	// iterate through players and dealer twice and deal a card on each iteration
	// similar to how a real game would be dealt
	for i := 0; i < 2; i++ {
		for _, v := range HandArr {
			card := deck[0]
			deck = deck[1:]
			v.hand = append(v.hand, card)
		}
	}

	return HandArr, deck
}

func drawCard(deck Deck) (Card, Deck) {
	card := deck[0]
	deck = deck[1:]
	return card, deck
}
