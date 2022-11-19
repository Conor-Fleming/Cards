package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Conor-Fleming/deck/deck"
)

type Card = deck.Card

type Hand struct {
	hand   []Card
	dealer bool
	bust   bool
	wins   int
	losses int
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

func (h Hand) checkBlackjack() bool {
	total, _ := getTotal(h)
	if total == 21 {
		fmt.Println("Blackjack!")
		fmt.Println()
		return true
	}
	return false
}

func (h Hand) checkBust() bool {
	total, _ := getTotal(h)
	if total > 21 {
		fmt.Println("Bust!")
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

// will need to change to accept array of players
// when adding bets can be checked here, element in struct?
func findWinner(dealer *Hand, players []*Hand) {
	dealerScore, _ := getTotal(*dealer)
	fmt.Println("Dealer -->", dealerScore)
	for i, player := range players[:len(players)-1] {
		//check if dealer was bust, if true and player !bust then player wins
		playerScore, _ := getTotal(*player)
		if !player.bust {
			var result string
			switch {
			case dealer.bust:
				result = "Win!"
				player.wins++
			case dealerScore > playerScore:
				result = "Lose."
				player.losses++
			case dealerScore < playerScore:
				result = "Win!"
				player.wins++
			default:
				result = "Push"
			}
			fmt.Printf("Player %v -> %v: %v\n", i+1, playerScore, result)
		} else {
			player.losses++
			fmt.Printf("Player %v -> %v: BUST\n", i+1, playerScore)
		}
	}
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

func stats(players []*Hand) string {
	statsOutput := "*** STATS ***\n"
	for i, player := range players[:len(players)-1] {
		statsOutput += fmt.Sprintf("Player %v:\nwins:%v   losses:%v\n", i+1, player.wins, player.losses)
	}
	return statsOutput
}

func printHit(card Card, player string) {
	fmt.Printf("%v hits.", player)
	time.Sleep(time.Second / 2)
	fmt.Print(".")
	time.Sleep(time.Second / 2)
	fmt.Print(".")
	time.Sleep(time.Second / 2)
	fmt.Println(card.String())
}
