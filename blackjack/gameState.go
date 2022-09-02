package main

import (
	"fmt"

	"github.com/omarahm3/gogo/deck/deck"
)

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("Its not player's turn yet")
	}
}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}

	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)

	return ret
}

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(3), deck.Shuffle)

	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 10)
	ret.Dealer = make(Hand, 0, 10)

	var card deck.Card

	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		addCard(&ret.Player, card)
		card, ret.Deck = draw(ret.Deck)
		addCard(&ret.Dealer, card)
	}

	ret.State = StatePlayerTurn

	return ret
}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()

	var card deck.Card

	card, ret.Deck = draw(ret.Deck)
	addCard(hand, card)

	if hand.Score() > 21 {
		return Stand(gs)
	}

	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)

	// Incrementing the state here is needed since states are ordered on the declaration "iota"
	ret.State++

	// This needed to be tested, or use a switch case as a more safe option

	return ret
}

func EndHand(gs GameState) GameState {
	ret := clone(gs)

	// Reveal all cards
	fmt.Println("Player Cards: ", ret.Player)
	fmt.Println("Dealer Cards: ", ret.Dealer)

	pScore := ret.Player.Score()
	dScore := ret.Dealer.Score()

	fmt.Println("------- Final Scores -------")
	fmt.Printf("Player: %d, Dealer: %d\n", pScore, dScore)

	switch {
	case pScore > MAX_SCORE:
		fmt.Println("You lose")
	case dScore > MAX_SCORE:
		fmt.Println("Dealer lost")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lose")
	case dScore == pScore:
		fmt.Println("Draw")
	}

	ret.Player = nil
	ret.Dealer = nil

	return ret
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func addCard(hand *Hand, card deck.Card) {
	*hand = append(*hand, card)
}
