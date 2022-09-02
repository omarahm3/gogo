package blackjackai

import (
	"fmt"

	"github.com/omarahm3/gogo/deck/deck"
)

const BET_AMOUNT = 10

type AI interface {
	Result(hand [][]deck.Card, dealer []deck.Card)
	Play(hand []deck.Card, dealer deck.Card) Move
	Bet(shuffled bool) int
}

func HumanAI() AI {
	return humanAI{}
}

type humanAI struct{}

func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player: ", hand)
		fmt.Println("Dealer: ", dealer)
		fmt.Print("What will you do? (h)it, (s)tand ")

		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		default:
			fmt.Println("invalid option:", input)
		}
	}
}

// [][]deck.Card because we will be supporting splitting, so player could potentially have multiple hands
func (ai humanAI) Result(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("Player Cards: ", hand)
	fmt.Println("Dealer Cards: ", dealer)
}

func (ai humanAI) Bet(shuffled bool) int {
	if shuffled {
		fmt.Println("The deck was just shuffled")
	}

	fmt.Print("What would you like to bet? ")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

type dealerAI struct{}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	if dealerCanDraw(hand) {
		return MoveHit
	}

	return MoveStand
}

func (ai dealerAI) Bet(shuffled bool) int {
	return BET_AMOUNT
}

func (ai dealerAI) Result(hand [][]deck.Card, dealer []deck.Card) {}
