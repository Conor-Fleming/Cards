// Package deck provides an easy implementation of a standard deck of cards.
// The package also includes some options such as shuffling, adding jokers, filtering cards, and adding multiple decks
package deck

import (
	"fmt"
	"math/rand"
	"time"
)

//go:generate stringer -type=Suit,Value

// The Suit type represents a cards suit and can be a Spade, Diamond, Club, Heart, or Joker
type Suit uint8

// A card will contain one of these suits
const (
	Spade Suit = iota + 1
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// The Value type represents a cards value Ace - King
type Value uint8

// A card will contain one of these values unless the card is of Suit Joker
const (
	Ace Value = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minVal = Ace  // minimum card Value
	maxVal = King // maximum card Value
)

// The Card type Reprents a Card from a deck of cards and contains a Value and a Suit
type Card struct {
	Value
	Suit
}

// String handles the printing of a cards Value and Suit
// If the card is of Suit Joker, the the String func will simply print "Joker"
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%v of %vs", c.Value.String(), c.Suit.String())
}

// New
func New(opts ...func([]Card) []Card) []Card {
	var deck []Card
	for _, suit := range suits {
		for val := minVal; val <= maxVal; val++ {
			card := Card{
				Value: val,
				Suit:  suit,
			}
			deck = append(deck, card)
		}
	}
	for _, opt := range opts {
		deck = opt(deck)
	}
	return deck
}

//func Sort(deck []Card) []Card {

//}

func Shuffle(deck []Card) []Card {
	temp := make([]Card, len(deck))
	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(len(deck))

	for i, v := range perm {
		temp[v] = deck[i]
	}
	return temp
}

func Jokers(n int) func([]Card) []Card {
	return func(c []Card) []Card {
		for i := 0; i < n; i++ {
			c = append(c, Card{
				Value: Value(i),
				Suit:  Joker,
			})
		}
		return c
	}
}

func Filter(filter func(card Card) bool) func([]Card) []Card {
	return func(c []Card) []Card {
		var filtered_slice []Card
		for _, val := range c {
			if !filter(val) {
				filtered_slice = append(filtered_slice, val)
			}
		}
		return filtered_slice
	}
}

func ExtraDecks(n int) func([]Card) []Card {
	return func(c []Card) []Card {
		extraDeck := c
		for i := 0; i < n; i++ {
			c = append(c, extraDeck...)
		}
		return c
	}
}
