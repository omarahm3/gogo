package blackjackai

import "testing"

func TestHit(t *testing.T) {
	game := New(Options{Hands: 1, Decks: 1})
	shuffle(&game)
	deal(&game)

	if game.state != statePlayerTurn {
		t.Errorf("expected game state to be player's turn, got %d", game.state)
	}

	MoveHit(&game)

	expectedLength := 3
	currentLength := len(game.player)

	if currentLength != expectedLength {
		t.Errorf("expected number of cards to be: %d, got: %d", expectedLength, currentLength)
	}
}
