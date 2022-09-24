package deck

//go:generate stringer -type=Suit,Value

type Suit uint8

const (
	Heart Suit = iota + 1
	Diamond
	Club
	Spade
	//Joker
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
	for _, suit := range _Suit_index {
		for _, val := range _Value_index {
			card := Card{
				Value: Value(val),
				Suit:  Suit(suit),
			}
			deck = append(deck, card)
		}
	}
	return deck
}
