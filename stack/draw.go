package stack

import (
	"errors"

	"github.com/TylerA73/go-solitaire/cards"
)

const (
	MAX = 3
)

type Draw struct {
	Cards []*cards.Card
}

func NewDraw() *Draw {
	return &Draw{
		Cards: make([]*cards.Card, 0),
	}
}

func (d *Draw) size() int {
	// since we are going to keep the slice at a static size, we have to
	// count the number of cards in the draw pile ourselves
	count := 0
	for i := 0; i < len(d.Cards); i++ {
		if d.Cards[i] != nil {
			count++
		}
	}
	return count
}

func (d *Draw) Draw(deck *Deck) {
	// If no cards left in the deck, do nothing
	if len(deck.Cards) == 0 {
		return
	}

	// Draw up to 3 new cards, and place the currect cards at the bottom of the deck

}

func (d *Draw) Pick() (*cards.Card, error) {
	// We can't pick from a draw pile that is empty
	if d.size() == 0 {
		return nil, errors.New("cannot pick from an empty draw pile. Please draw from the deck first")
	}

	// Pick the last card from the draw pile
	var card *cards.Card
	for i := d.size() - 1; i > 0; i-- {
		if d.Cards[i] != nil {
			card = d.Cards[i]
			d.Cards[i] = nil
			break
		}
	}

	return card, nil
}
