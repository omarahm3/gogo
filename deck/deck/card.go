//go:generate stringer -type=Rank,Suit
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8
type Rank uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Rank
	Suit
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(opts ...func([]Card) []Card) []Card {
	var deck []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			deck = append(deck, Card{
				Suit: suit,
				Rank: rank,
			})
		}
	}

	for _, opt := range opts {
		deck = opt(deck)
	}

	return deck
}

func DefaultSort(deck []Card) []Card {
	sort.Slice(deck, Less(deck))
	return deck
}

func Less(deck []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absoluteRank(deck[i]) < absoluteRank(deck[j])
	}
}

func Sort(less func(deck []Card) func(i, j int) bool) func([]Card) []Card {
	return func(deck []Card) []Card {
		sort.Slice(deck, less(deck))
		return deck
	}
}

func Shuffle(deck []Card) []Card {
	ret := make([]Card, len(deck))
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i, j := range r.Perm(len(deck)) {
		ret[i] = deck[j]
	}

	return ret
}

func Jokers(n int) func([]Card) []Card {
	return func(deck []Card) []Card {
		for i := 0; i < n; i++ {
			deck = append(deck, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}

		return deck
	}
}

func absoluteRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
