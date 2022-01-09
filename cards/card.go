package cards

import (
	"github.com/TylerA73/go-solitaire/suits"
)

type Card struct {
	Ordinal   int
	Suit      *suits.Suit
	IsFlipped bool
}

func NewCard(ordinal int, suit *suits.Suit) *Card {
	return &Card{
		Ordinal:   ordinal,
		Suit:      suit,
		IsFlipped: false,
	}
}

func (c *Card) IsRed() bool {
	return c.Suit.IsRed
}

func (c *Card) Flip() {
	c.IsFlipped = !c.IsFlipped
}
