package deck

import (
	"reflect"
	"testing"
)

var deck = []Card{
	{Ace, Spade}, {Two, Spade}, {Three, Spade}, {Four, Spade}, {Five, Spade}, {Six, Spade}, {Seven, Spade}, {Eight, Spade}, {Nine, Spade}, {Ten, Spade}, {Jack, Spade}, {Queen, Spade}, {King, Spade},
	{Ace, Diamond}, {Two, Diamond}, {Three, Diamond}, {Four, Diamond}, {Five, Diamond}, {Six, Diamond}, {Seven, Diamond}, {Eight, Diamond}, {Nine, Diamond}, {Ten, Diamond}, {Jack, Diamond}, {Queen, Diamond}, {King, Diamond},
	{Ace, Club}, {Two, Club}, {Three, Club}, {Four, Club}, {Five, Club}, {Six, Club}, {Seven, Club}, {Eight, Club}, {Nine, Club}, {Ten, Club}, {Jack, Club}, {Queen, Club}, {King, Club},
	{Ace, Heart}, {Two, Heart}, {Three, Heart}, {Four, Heart}, {Five, Heart}, {Six, Heart}, {Seven, Heart}, {Eight, Heart}, {Nine, Heart}, {Ten, Heart}, {Jack, Heart}, {Queen, Heart}, {King, Heart},
}

func TestNew(t *testing.T) {
	expected := deck
	result := New()
	equalityCheck := reflect.DeepEqual(expected, result)
	if equalityCheck != true {
		t.Errorf("got %q, wanted %q", result, expected)
	}
}

func TestShuffle(t *testing.T) {
	first := Shuffle(deck)
	second := Shuffle(deck)
	isEqual := reflect.DeepEqual(first, second)
	if isEqual == true {
		t.Errorf("Error, both versions of the shuffled deck are the same\n %v,\n %v", first, second)
	}
}

func TestJokers(t *testing.T) {
	for i := 1; i < 5; i++ {
		one_joker := Jokers(i)
		newDeck := New(one_joker)
		if len(newDeck) != (52 + i) {
			t.Errorf("wrong amount of Jokers in deck. Expected %v jokers but have %v", i, len(newDeck)-i)
		}
	}
}

func TestExtraDecks(t *testing.T) {
	for i := 1; i < 4; i++ {
		testDeck := New(ExtraDecks(i))
		if len(testDeck) != (52 + (i * 52)) {
			t.Errorf("wrong amount of cards in this deck have %v,\nwant %v", len(testDeck), 52+(i*52))
		}
	}
}

func TestFilter(t *testing.T) {
	filt := func(card Card) bool {
		return card.Value == 2 || card.Value == 5 || card.Value == 10
	}
	deck := New(Filter(filt))
	for _, v := range deck {
		if v.Value == 2 || v.Value == 5 || v.Value == 10 {
			t.Errorf("error, %v not expected", v.Value)
		}
	}
}
