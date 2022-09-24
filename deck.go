package deck

//go:generate stringer -type=Suit,Value

type Suit uint8

const (
	Heart Suit = iota + 1
	Diamond
	Club
	Spade
)

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

type Card struct {
	Value
	Suit
}

func (c *Card) New() []Card {
	var deck []Card
	for i := 1; i < 5; i++ {
		for y := 1; y < 14; y++ {
			card := Card{
				Value: Value(y),
				Suit:  Suit(i),
			}
			deck = append(deck, card)
		}
	}
	return deck

}
