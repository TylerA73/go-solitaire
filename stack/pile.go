package stack

import "github.com/TylerA73/go-solitaire/cards"

type Pile struct {
	Stack
}

func NewPile() *Pile {
	return &Pile{
		Stack{
			Cards: make([]cards.Card, 0),
		},
	}
}

func (p *Pile) CanStack(card *cards.Card) bool {
	// For piles, we don't care about suit
	// If pile is empty, just allow anything to be stacked onto it
	// If it's not empty, then the colour of the suits of the top card and
	// the stacking card must not be the same, and the stacking card has to be
	// equal to the last card - 1
	var topCard cards.Card
	if len(p.Cards) == 0 {
		return true
	} else {
		topCard = p.Cards[len(p.Cards)-1]
	}
	return topCard.Ordinal == card.Ordinal+1 && topCard.IsRed() != card.IsRed()
}

func (p *Pile) Add(card cards.Card) {
	p.Cards = append(p.Cards, card)
}
