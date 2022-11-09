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

// version 1
func main() {
	//Decks sets the amount of extra decks to add to the intial deck
	deck := deck.New(deck.ExtraDecks(2), deck.Shuffle)

	// get input for how many players
	// ?

	// call deal function with that amount
	allPlayers, deck := deal(deck, 1)
	for i, v := range allPlayers {
		if i == len(allPlayers)-1 {
			continue
		}
		total := getTotal(v.hand)
		fmt.Printf("Player %v: %v  Total-----> %v\n", i+1, v.String(), total)
	}
	fmt.Println("Dealer:", allPlayers[len(allPlayers)-1].String())

	var move string

	for i := 0; i < len(allPlayers)-1; i++ {
		for {
			fmt.Println("Hit or stand?")
			fmt.Scanf("%s", &move)
			move = strings.ToLower(move)

			// getting same value from each draw call
			if move == "hit" {
				hit, _ := drawCard(deck)
				fmt.Println(hit.String())
				allPlayers[i].hand = append(allPlayers[i].hand, hit)
				fmt.Println(allPlayers[i].String())
				continue
			}

			if move == "stand" {
				break
			}

			fmt.Println("Invalid entry")
		}
	}
}

func getTotal(hand []deck.Card) int {
	var total int
	for _, v := range hand {
		total += int(v.Value)
	}
	return total
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
