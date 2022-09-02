package main

import (
	"strings"

	"github.com/omarahm3/gogo/deck/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))

	for i := range strs {
		strs[i] = h[i].String()
	}

	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func (h Hand) MinimumScore() int {
	score := 0

	for _, c := range h {
		// Since jack, queen and king are all worth of 10 points, minimum is used here
		score += minimum(int(c.Rank), 10)
	}

	return score
}

func (h Hand) Score() int {
	minScore := h.MinimumScore()

	// we no longer need to count aces as 11 anymore
	if minScore > 11 {
		return minScore
	}

	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}

	return minScore
}

func (h Hand) DealerCanDraw() bool {
	return h.Score() <= 16 || (h.Score() == 17 && h.MinimumScore() != 17)
}

func minimum(a, b int) int {
	if a < b {
		return a
	}

	return b
}
