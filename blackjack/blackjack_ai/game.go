package blackjackai

import (
	"errors"

	"github.com/omarahm3/gogo/deck/deck"
)

const (
	MAX_SCORE                = 21
	DEFAULT_NUMBER_OF_DECKS  = 3
	DEFAULT_NUMBER_OF_HANDS  = 100
	DEFAULT_BLACKJACK_PAYOUT = 1.5
)

var (
	errBust             = errors.New("hand score exceeded 21")
	errLessThanTwoCards = errors.New("hand has less than 2 cards")
	errInvalidState     = errors.New("invalid state")
)

type state int8

type Move func(*Game) error

const (
	stateBet state = iota
	statePlayerTurn
	stateDealerTurn
	stateHandOver
)

type Options struct {
	BlackJackPayout float64
	Decks           int
	Hands           int
}

type hand struct {
	cards []deck.Card
	bet   int
}

type Game struct {
	nDecks          int
	nHands          int
	blackJackPayout float64

	deck  []deck.Card
	state state

	player    []hand
	handIndex int
	playerBet int
	balance   int

	dealer   []deck.Card
	dealerAI AI
}

func New(opts Options) Game {
	g := Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
	}

	if opts.Hands == 0 {
		opts.Hands = DEFAULT_NUMBER_OF_HANDS
	}

	if opts.Decks == 0 {
		opts.Decks = DEFAULT_NUMBER_OF_DECKS
	}

	if opts.BlackJackPayout == 0 {
		opts.BlackJackPayout = DEFAULT_BLACKJACK_PAYOUT
	}

	g.nDecks = opts.Decks
	g.nHands = opts.Hands
	g.blackJackPayout = opts.BlackJackPayout

	return g
}

func (g *Game) currentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player[g.handIndex].cards
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("Its not player's turn yet")
	}
}

func (g *Game) Play(ai AI) int {
	g.deck = nil
	minNumOfCardsTillShuffle := 52 * g.nDecks / 3

	for i := 0; i < g.nHands; i++ {
		shuffled := false
		if len(g.deck) < minNumOfCardsTillShuffle {
			shuffle(g)
			shuffled = true
		}

		bet(g, ai, shuffled)
		deal(g)

		if Blackjack(g.dealer...) {
			endRound(g, ai)
			continue
		}

		for g.state == statePlayerTurn {
			// So that in case the ai.Play is modifying the hand
			// We copy this so that it won't modify our game state
			hand := make([]deck.Card, len(*g.currentHand()))
			copy(hand, *g.currentHand())
			move := ai.Play(hand, g.dealer[0])
			err := move(g)

			switch err {
			case errBust:
				MoveStand(g)
			case nil:
				// do nothing
			default:
				panic(err)
			}
		}

		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			// second argument doesn't really mean anything, just for consistency it will be the first card of dealer's han
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}

		endRound(g, ai)
	}

	return g.balance
}

func MoveHit(g *Game) error {
	hand := g.currentHand()

	var card deck.Card

	card, g.deck = draw(g.deck)
	addCard(hand, card)

	if Score(*hand...) > 21 {
		return errBust
	}
	return nil
}

func MoveStand(g *Game) error {
	if g.state == stateDealerTurn {
		g.state++
		return nil
	}

	if g.state == statePlayerTurn {
		g.handIndex++
		if g.handIndex >= len(g.player) {
			// Incrementing the state here is needed since states are ordered on the declaration "iota"
			g.state++
		}
		return nil
	}

	return errInvalidState
}

func MoveDouble(g *Game) error {
	if len(g.player) != 2 {
		return errLessThanTwoCards
	}

	g.playerBet *= 2
	MoveHit(g)
	return MoveStand(g)
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

func Blackjack(hand ...deck.Card) bool {
	return len(hand) == 2 && Score(hand...) == MAX_SCORE
}

func shuffle(g *Game) {
	g.deck = deck.New(deck.Deck(g.nDecks), deck.Shuffle)
}

func deal(g *Game) {
	playerHand := make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)

	var card deck.Card

	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		addCard(&playerHand, card)
		card, g.deck = draw(g.deck)
		addCard(&g.dealer, card)
	}

	g.player = []hand{
		{
			cards: playerHand,
			bet:   g.playerBet,
		},
	}
	g.state = statePlayerTurn
}

func endRound(g *Game, ai AI) {
	dScore := Score(g.dealer...)
	dBlackjack := Blackjack(g.dealer...)
	allHands := make([][]deck.Card, len(g.player))

	for i, hand := range g.player {
		cards := hand.cards
		allHands[i] = cards
		winnings := hand.bet
		pScore, pBlackjack := Score(cards...), Blackjack(cards...)

		switch {
		case pBlackjack && dBlackjack:
			winnings = 0
		case dBlackjack:
			winnings *= -1
		case pBlackjack:
			winnings *= int(g.blackJackPayout)
		case pScore > MAX_SCORE:
			winnings *= -1
		case dScore > MAX_SCORE:
			// win
		case pScore > dScore:
			// win
		case dScore > pScore:
			winnings *= -1
			winnings = 0
		}

		g.balance += winnings
	}

	ai.Result(allHands, g.dealer)

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

func bet(g *Game, ai AI, shuffled bool) {
	g.playerBet = ai.Bet(shuffled)
}

func minimum(a, b int) int {
	if a < b {
		return a
	}

	return b
}
