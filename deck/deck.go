package deck

import (
	"math/rand"
	"sort"
	"time"
)

type CardValue int

const (
	V_Joker CardValue = iota
	V_A
	V_1
	V_2
	V_3
	V_4
	V_5
	V_6
	V_7
	V_8
	V_9
	V_10
	V_J
	V_Q
	V_K
)

type CardSuit int

const (
	S_Joker CardSuit = iota
	S_Spades
	S_Diamonds
	S_Clubs
	S_Hearts
)

type Card struct {
	Value CardValue
	Suit  CardSuit
}

func (c *Card) Equals(card *Card) bool {
	return c.Value == card.Value && c.Suit == card.Suit
}

type Deck []Card

func New() Deck {
	result := make(Deck, 0)

	for suit := CardSuit(1); suit < 5; suit++ {
		for value := CardValue(1); value < 14; value++ {
			result = append(result, Card{value, suit})
		}
	}

	return result
}

func NewComposed(deckAmount int) Deck {
	result := make(Deck, 0)

	for i := 0; i < deckAmount; i++ {
		result = append(result, New()...)
	}

	return result
}

func (d Deck) Len() int {
	return len(d)
}

func (d Deck) Less(i, j int) bool {
	return int(d[i].Suit)*14+int(d[i].Value) <
		int(d[j].Suit)*14+int(d[j].Value)
}

func (d Deck) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Deck) AddJoker() Deck {
	return append(d, Card{V_Joker, S_Joker})
}

func (d Deck) AddJokers(amount int) Deck {
	for i := 0; i < amount; i++ {
		d = d.AddJoker()
	}

	return d
}

func (d Deck) ShuffleWithSeed(seed int64) Deck {
	rand.Seed(seed)
	rand.Shuffle(d.Len(), d.Swap)
	return d
}

func (d Deck) Shuffle() Deck {
	d.ShuffleWithSeed(time.Now().UnixNano())
	return d
}

func (d Deck) SortDefault() Deck {
	sort.Sort(d)
	return d
}

func (d Deck) Sort(less func(int, int) bool) {
	sort.SliceStable(d, less)
}

func (d Deck) Filter(leave func(Card) bool) Deck {
	filtered := make(Deck, 0)

	for i := range d {
		if leave(d[i]) {
			filtered = append(filtered, d[i])
		}
	}

	return filtered
}

func (d Deck) HasCard(card Card) bool {
	for _, c := range d {
		if c.Value == card.Value && c.Suit == card.Suit {
			return true
		}
	}

	return false
}
