package main

import (
	"fmt"
)

const MAX_SCORE = 21

func main() {
	var gs GameState
	gs = Shuffle(gs)
	gs = Deal(gs)

	var input string

	for gs.State == StatePlayerTurn {
		fmt.Println("Player: ", gs.Player)
		fmt.Println("Dealer: ", gs.Dealer.DealerString())
		fmt.Println("What will you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			gs = Hit(gs)
		case "s":
			gs = Stand(gs)
		default:
			fmt.Println("invalid option:", input)
		}
	}

	for gs.State == StateDealerTurn {
		if gs.Dealer.DealerCanDraw() {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}

  EndHand(gs)
}
