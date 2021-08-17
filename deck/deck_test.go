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
	if !first.Equals(&deck[deck.Len()-1]) {
		t.Error("Swap doesn't swaps")
	}
}

func TestShufflingTheDeck(t *testing.T) {
	deck := New()

	deck.Shuffle()

	if sort.SliceIsSorted(deck, deck.Less) {
		t.Error("Shuffle() doesn't shuffles the deck.")
	}

	shuffledDeck := make(Deck, len(deck))
	copy(shuffledDeck, deck)

	deck.Shuffle()

	different := false
	for i, card := range deck {
		if card.Value != shuffledDeck[i].Value || card.Suit != shuffledDeck[i].Suit {
			different = true
			break
		}
	}

	if !different {
		t.Error("Shuffle() is not randomized.")
	}
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

	if !deck.HasCard(Card{V_Joker, S_Joker}) {
		t.Error("Adding Jokers failed")
	}
}

func TestDeckFiltering(t *testing.T) {
	type test_t struct {
		values []CardValue
		suits  []CardSuit
		cards  []Card
	}

	tests := []test_t{
		{
			[]CardValue{},
			[]CardSuit{},
			[]Card{{V_2, S_Hearts}},
		}, {
			[]CardValue{V_2, V_3},
			[]CardSuit{S_Diamonds},
			[]Card{{V_J, S_Hearts}},
		}, {
			[]CardValue{},
			[]CardSuit{S_Clubs},
			[]Card{},
		},
	}

	for _, test := range tests {
		testname := fmt.Sprint(test.values, test.suits, test.cards)
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
					if card.Equals(&c) {
						return false
					}
				}

				return true
			}

			deck.Filter(filter)

			for _, card := range deck {
				if !filter(card) {
					t.Errorf("Filter is not working: card %s should has been deleted.\nDeck: %s", card, deck)
					break
				}
			}

		})

	}
}
