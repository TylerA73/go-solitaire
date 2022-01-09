package stack

import (
	"github.com/TylerA73/go-solitaire/cards"
	"github.com/TylerA73/go-solitaire/suits"
)

type Home struct {
	Stack
	suit *suits.Suit
}

func NewHome(suit *suits.Suit) *Home {
	return &Home{
		Stack: Stack{
			Cards: make([]cards.Card, 0),
		},
		suit: suit,
	}
}

func (h *Home) CanStack(card *cards.Card) bool {
	// If suits don't match, then the card can't be stacked
	if !h.suit.IsEqual(card.Suit) {
		return false
	}

	// Card suit must be the same as the home pile suit
	// If card is an ace, and nothing is on the pile, we can stack it
	// otherwise the card being stacked has to be equal to the top card + 1
	var topCard cards.Card
	if len(h.Cards) == 0 && card.Ordinal == 1 {
		return true
	} else {
		topCard = h.Cards[len(h.Cards)-1]
	}

	return topCard.Ordinal+1 == card.Ordinal && topCard.Suit.IsEqual(card.Suit)
}

func (h *Home) Add(card cards.Card) {
	h.Cards = append(h.Cards, card)
}
