package stack

import "github.com/TylerA73/go-solitaire/cards"

type Deck struct {
	Stack
}

func (d *Deck) Pick() *cards.Card {
	return nil
}

func (d *Deck) Shuffle() {

}
