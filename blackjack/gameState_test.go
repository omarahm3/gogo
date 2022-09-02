package main

import "testing"

func TestHit(t *testing.T) {
	var gs GameState
	gs = Shuffle(gs)
	gs = Deal(gs)

	if gs.State != StatePlayerTurn {
		t.Errorf("expected game state to be player's turn, got %d", gs.State)
	}

	gs = Hit(gs)

	expectedLength := 3
	currentLength := len(gs.Player)

	if currentLength != expectedLength {
		t.Errorf("expected number of cards to be: %d, got: %d", expectedLength, currentLength)
	}
}
