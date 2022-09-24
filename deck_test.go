package deck

import (
	"reflect"
	"testing"
)

var deck = []Card{
	{Ace, Heart}, {Two, Heart}, {Three, Heart}, {Four, Heart}, {Five, Heart}, {Six, Heart}, {Seven, Heart}, {Eight, Heart}, {Nine, Heart}, {Ten, Heart}, {Jack, Heart}, {Queen, Heart}, {King, Heart},
	{Ace, Diamond}, {Two, Diamond}, {Three, Diamond}, {Four, Diamond}, {Five, Diamond}, {Six, Diamond}, {Seven, Diamond}, {Eight, Diamond}, {Nine, Diamond}, {Ten, Diamond}, {Jack, Diamond}, {Queen, Diamond}, {King, Diamond},
	{Ace, Club}, {Two, Club}, {Three, Club}, {Four, Club}, {Five, Club}, {Six, Club}, {Seven, Club}, {Eight, Club}, {Nine, Club}, {Ten, Club}, {Jack, Club}, {Queen, Club}, {King, Club},
	{Ace, Spade}, {Two, Spade}, {Three, Spade}, {Four, Spade}, {Five, Spade}, {Six, Spade}, {Seven, Spade}, {Eight, Spade}, {Nine, Spade}, {Ten, Spade}, {Jack, Spade}, {Queen, Spade}, {King, Spade},
}

func TestNew(t *testing.T) {
	c := Card{}
	expected := deck
	result := c.New()
	equalityCheck := reflect.DeepEqual(expected, result)
	if equalityCheck != true {
		t.Errorf("got %q, wanted %q", result, expected)
	}
}
