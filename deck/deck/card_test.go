package deck

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{
		Rank: Ace,
		Suit: Heart,
	})

	fmt.Println(Card{
		Suit: Joker,
	})

	// Output:
	// Ace of Hearts
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	expectedDeckSize := 13 * 4

	if len(cards) != expectedDeckSize {
		t.Errorf("expected %d, go %d", expectedDeckSize, len(cards))
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	firstCard := Card{Rank: Ace, Suit: Spade}

	if cards[0] != firstCard {
		t.Errorf("expected first card to be: %q, got %q", firstCard, cards[0])
	}
}

func TestShuffle(t *testing.T) {
	defaultCards := New()
	cards := New(Shuffle)

	if reflect.DeepEqual(defaultCards, cards) {
		t.Error("cards are not shuffled")
	}
}
