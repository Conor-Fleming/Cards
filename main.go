package main

import (
	"fmt"

	"github.com/Conor-Fleming/deck/deck"
)

func main() {
	deck := deck.New(deck.Shuffle)
	fmt.Println(deck)
}
