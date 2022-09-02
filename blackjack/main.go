package main

import (
	"fmt"

	blackjackai "github.com/omarahm3/gogo/blackjack/blackjack_ai"
)

func main() {
  game := blackjackai.New()
  winnings := game.Play(blackjackai.HumanAI())

  fmt.Println(winnings)
}
