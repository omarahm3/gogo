package main

import (
	"fmt"

	blackjackai "github.com/omarahm3/gogo/blackjack/blackjack_ai"
)

func main() {
  game := blackjackai.New(blackjackai.Options{Hands: 1})
  winnings := game.Play(blackjackai.HumanAI())

  fmt.Println(winnings)
}
