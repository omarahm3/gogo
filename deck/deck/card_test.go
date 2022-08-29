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

func TestJookers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0

	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}

	if count != 3 {
		t.Errorf("expected 3 jokers, got %d", count)
	}
}

func TestFilter(t *testing.T) {
	cards := New(Filter(func(card Card) bool {
		return card.Rank == Ace && card.Suit == Spade
	}))

	for _, c := range cards {
		if c.Rank == Ace && c.Suit == Spade {
			t.Errorf("%s is not filtered out of cards", c)
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	count := 0

	for _, c := range cards {
		if c.Rank == Ace {
			count++
		}
	}

	if count != 12 {
		t.Errorf("expected to have 12 Aces, got: %d", count)
	}
}
