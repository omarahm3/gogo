package blackjackai

import "testing"

func TestHit(t *testing.T) {
	var gs Game
	gs = shuffle(gs)
	gs = deal(gs)

	if gs.state != statePlayerTurn {
		t.Errorf("expected game state to be player's turn, got %d", gs.state)
	}

	gs = MoveHit(gs)

	expectedLength := 3
	currentLength := len(gs.player)

	if currentLength != expectedLength {
		t.Errorf("expected number of cards to be: %d, got: %d", expectedLength, currentLength)
	}
}
