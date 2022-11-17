# Blackjack
[![GgoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/Conor-Fleming/blackjack/deck)
[![Go Report Card](https://goreportcard.com/badge/github.com/Conor-Fleming/deck)](https://goreportcard.com/report/github.com/Conor-Fleming/deck)

A command line interface blackjack game

### Playing the game
To play this game simply clone this repository 
```git clone https://github.com/Conor-Fleming/blackjack.git```

navigate to the new directory and run `./blackjack` to play.

### Features
The blackjack program allows the user(s) to specify the number of players at the beginning of the game. This allows you to go 1:1 with the dealer or enjoy the game with any/all of your friends.

This blackjack game uses a custom Deck package that allows for customizations of decks and other things such as:
- multiple decks
- filtering cards
- jokers
- sorting

To read more about the deck package please see [here](http://godoc.org/github.com/Conor-Fleming/blackjack/deck).

### Roadmap
My future plans for this project would be to add support for some sort of betting system and keeping track of wins and losses.
