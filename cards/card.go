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

func (c *Card) CanStackOnPile(card *Card) bool {
	return c.Ordinal == card.Ordinal+1 && c.IsRed() != card.IsRed()
}

func (c *Card) CanStackOnHome(card *Card) bool {
	return c.Ordinal == card.Ordinal-1 && c.Suit.IsEqual(card.Suit)
}

func (c *Card) Flip() {
	c.IsFlipped = !c.IsFlipped
}
