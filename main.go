package main

import (
	"fmt"

	"github.com/Conor-Fleming/deck/deck"
)

type Card = deck.Card

type Hand []deck.Card

type players struct {
	dealer bool
	hand   Hand
}

// version 1
func main() {
	//Decks sets the amount of extra decks to add to the intial deck
	deck := deck.New(deck.ExtraDecks(2), deck.Shuffle)

	fmt.Println(deck[0])
	deck = deck[1:]
	fmt.Println(deck[0])

	//fmt.Println("The dealers hand:\n", dealer[1:])
	//fmt.Println("Player 1's hand:\n", player)

	//player can hit or stand
	//var hit string
	//fmt.Println("Press 'h' to hit or 's' to stay")
	//fmt.Scan(&hit)
	//if hit == "h" {
	//player1.Hand = append(player1.Hand, playerHit(deck))
	//}
	//fmt.Println("Player 1's hand:\n", player1.Hand)

}

// will always be only 1 player for first version of program
func deal(players int) []players

func playerHit(deck []Card) Card {
	card := deck[0]
	deck = deck[1:]
	return card
}
