package deck

import (
	"fmt"
	"sort"
	"testing"
)

func TestCreatedDeckIsNotEmpty(t *testing.T) {
	deck := New()
	if deck.Len() == 0 {
		t.Error("Empty Deck was created")
	}
}

func TestCardSwap(t *testing.T) {
	deck := New()
	first := deck[0]
	deck.Swap(0, deck.Len()-1)
	if first != deck[deck.Len()-1] {
		t.Error("Swap doesn't swaps")
	}
}

func TestShufflingTheDeck(t *testing.T) {
	deck := New()

	deck.Shuffle()

	if sort.SliceIsSorted(deck, deck.Less) {
		t.Error("Shuffle() doesn't shuffles the deck.")
	}

	newDeck := New()

	newDeck.Shuffle()

	for i, card := range deck {
		if card != newDeck[i] {
			return
		}
	}

	t.Error("Shuffle() is not randomized.")
}

func TestDefaultSort(t *testing.T) {
	deck := New()

	deck.Shuffle()
	deck.SortDefault()
	if !sort.IsSorted(deck) {
		t.Error("SortDefault() is not using deck.Less() function.")
	}
}

func TestAddingJoker(t *testing.T) {
	deck := New().AddJoker()

	if !deck.HasCard(Card{VJoker, SJoker}) {
		t.Error("Adding Jokers failed")
	}
}

func TestDeckFiltering(t *testing.T) {
	tests := []struct {
		values []CardValue
		suits  []CardSuit
		cards  []Card
	}{
		{
			[]CardValue{},
			[]CardSuit{},
			[]Card{{V2, SHearts}},
		}, {
			[]CardValue{V2, V3},
			[]CardSuit{SDiamonds},
			[]Card{{VJ, SHearts}},
		}, {
			[]CardValue{},
			[]CardSuit{SClubs},
			[]Card{},
		},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s %s %s", test.values, test.suits, test.cards)
		t.Run(testname, func(t *testing.T) {
			deck := New()

			filter := func(card Card) bool {
				for _, v := range test.values {
					if card.Value == v {
						return false
					}
				}

				for _, s := range test.suits {
					if card.Suit == s {
						return false
					}
				}

				for _, c := range test.cards {
					if card == c {
						return false
					}
				}

				return true
			}

			filteredDeck := deck.Filter(filter)

			for _, card := range filteredDeck {
				if !filter(card) {
					t.Fatalf("Filter is not working: card %s should has been deleted.\nDeck: %s", card, deck)
				}
			}

		})

	}
}
