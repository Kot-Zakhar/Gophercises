package deck

import (
	"math/rand"
	"sort"
	"time"
)

type CardValue int

const (
	VJoker CardValue = iota
	VA
	V1
	V2
	V3
	V4
	V5
	V6
	V7
	V8
	V9
	V10
	VJ
	VQ
	VK
	VAmount = VK
)

type CardSuit int

const (
	SJoker CardSuit = iota
	SSpades
	SDiamonds
	SClubs
	SHearts
	SAmount = SHearts
)

type Card struct {
	Value CardValue
	Suit  CardSuit
}

type Deck []Card

func New() Deck {
	result := make(Deck, 0)

	for suit := CardSuit(1); suit < SAmount; suit++ {
		for value := CardValue(1); value < VAmount; value++ {
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
	return append(d, Card{VJoker, SJoker})
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
