package deck

import "testing"

var deck = []Card{
	Card{Ace, Heart}, Card{Two, Heart}, Card{Three, Heart}, Card{Four, Heart}, Card{Five, Heart}, Card{Six, Heart}, Card{Seven, Heart}, Card{Eight, Heart}, Card{Nine, Heart}, Card{Ten, Heart}, Card{Jack, Heart}, Card{Queen, Heart}, Card{King, Heart},
	Card{Ace, Diamond}, Card{Two, Diamond}, Card{Three, Diamond}, Card{Four, Diamond}, Card{Five, Diamond}, Card{Six, Diamond}, Card{Seven, Diamond}, Card{Eight, Diamond}, Card{Nine, Diamond}, Card{Ten, Diamond}, Card{Jack, Diamond}, Card{Queen, Diamond}, Card{King, Diamond},
	Card{Ace, Club}, Card{Two, Club}, Card{Three, Club}, Card{Four, Club}, Card{Five, Club}, Card{Six, Club}, Card{Seven, Club}, Card{Eight, Club}, Card{Nine, Club}, Card{Ten, Club}, Card{Jack, Club}, Card{Queen, Club}, Card{King, Club},
	Card{Ace, Spade}, Card{Two, Spade}, Card{Three, Spade}, Card{Four, Spade}, Card{Five, Spade}, Card{Six, Spade}, Card{Seven, Spade}, Card{Eight, Spade}, Card{Nine, Spade}, Card{Ten, Spade}, Card{Jack, Spade}, Card{Queen, Spade}, Card{King, Spade},
}

func TestNew(t *testing.T) {
	expected := deck
	result := New()
	if result != expected {
		t.Errorf("got %q, wanted %q", result, expected)
	}

}
