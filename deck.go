package deck

import (
	"fmt"
	"math/rand"
	"time"
)

//go:generate stringer -type=Suit,Value

type Suit uint8

const (
	Spade Suit = iota + 1
	Diamond
	Club
	Heart
	//Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Value uint8

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
	minVal = Ace
	maxVal = King
)

type Card struct {
	Value
	Suit
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %vs", c.Value.String(), c.Suit.String())
}

// using functional options pattern to allow users to specify states of the Card deck i.e. sorted/shuffled/with jokers etc
// without any options this will return the deck in the form of a brand new deck each suit in ascending order (A-K)
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
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}
