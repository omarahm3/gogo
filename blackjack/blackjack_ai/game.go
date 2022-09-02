package blackjackai

import (
	"fmt"

	"github.com/omarahm3/gogo/deck/deck"
)

const (
	MAX_SCORE = 21
  NUMBER_OF_HANDS = 2
)

type state int8

type Move func(*Game)

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

type Game struct {
	deck     []deck.Card
	state    state
	player   []deck.Card
	dealer   []deck.Card
	dealerAI AI
	balance  int
}

func (g *Game) currentPlayer() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("Its not player's turn yet")
	}
}

func (g *Game) Play(ai AI) int {
	shuffle(g)
	for i := 0; i < NUMBER_OF_HANDS; i++ {
		deal(g)

		for g.state == statePlayerTurn {
			// So that in case the ai.Play is modifying the hand
			// We copy this so that it won't modify our game state
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := ai.Play(hand, g.dealer[0])
			move(g)
		}

		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			// second argument doesn't really mean anything, just for consistency it will be the first card of dealer's han
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}

		endHand(g, ai)
	}

	return g.balance
}

func New() Game {
	return Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
	}
}

func MoveHit(g *Game) {
	hand := g.currentPlayer()

	var card deck.Card

	card, g.deck = draw(g.deck)
	addCard(hand, card)

	if Score(*hand...) > 21 {
		MoveStand(g)
	}
}

func MoveStand(g *Game) {
	// Incrementing the state here is needed since states are ordered on the declaration "iota"
	g.state++
}

func Score(hand ...deck.Card) int {
	minScore := minimumScore(hand...)

	// we no longer need to count aces as 11 anymore
	if minScore > 11 {
		return minScore
	}

	for _, c := range hand {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}

	return minScore
}

// Returns true if score of a hand is a soft score - that is if an ace is being counted as 11 points
func Soft(hand ...deck.Card) bool {
	minScore := minimumScore(hand...)
	score := Score(hand...)

	return minScore != score
}

func shuffle(g *Game) {
	g.deck = deck.New(deck.Deck(3), deck.Shuffle)
}

func deal(g *Game) {
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)

	var card deck.Card

	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		addCard(&g.player, card)
		card, g.deck = draw(g.deck)
		addCard(&g.dealer, card)
	}

	g.state = statePlayerTurn
}

func endHand(g *Game, ai AI) {
	pScore, dScore := Score(g.player...), Score(g.dealer...)

	switch {
	case pScore > MAX_SCORE:
		fmt.Println("You lose")
		g.balance--
	case dScore > MAX_SCORE:
		fmt.Println("Dealer lost")
		g.balance++
	case pScore > dScore:
		fmt.Println("You win!")
		g.balance++
	case dScore > pScore:
		fmt.Println("You lose")
		g.balance--
	case dScore == pScore:
		fmt.Println("Draw")
	}

	ai.Result([][]deck.Card{g.player}, g.dealer)

	g.player = nil
	g.dealer = nil
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func addCard(hand *[]deck.Card, card deck.Card) {
	*hand = append(*hand, card)
}

func minimumScore(hand ...deck.Card) int {
	score := 0

	for _, c := range hand {
		// Since jack, queen and king are all worth of 10 points, minimum is used here
		score += minimum(int(c.Rank), 10)
	}

	return score
}

func dealerCanDraw(dealer []deck.Card) bool {
	dealerScore := Score(dealer...)
	return dealerScore <= 16 || (dealerScore == 17 && Soft(dealer...))
}

func minimum(a, b int) int {
	if a < b {
		return a
	}

	return b
}
