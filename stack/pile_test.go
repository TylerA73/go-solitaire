package stack

import (
	"testing"

	"github.com/TylerA73/go-solitaire/cards"
	"github.com/TylerA73/go-solitaire/suits"
	"github.com/stretchr/testify/assert"
)

func Test_PileCanStack(t *testing.T) {
	t.Run("should be able to stack card onto an empty pile", func(t *testing.T) {
		pile := NewPile()

		assert.True(t, pile.CanStack(card))
	})

	t.Run("should be able to stack card onto a pile with cards", func(t *testing.T) {
		pile := NewPile()
		pile.Add(*cards.NewCard(2, &suits.Diamonds))

		assert.True(t, pile.CanStack(card))
		assert.True(t, pile.CanStack(cards.NewCard(1, &suits.Spades)))
	})

	t.Run("should not be able to stack card onto a pile with cards", func(t *testing.T) {
		pile := NewPile()
		pile.Add(*cards.NewCard(2, &suits.Clubs))

		assert.False(t, pile.CanStack(card))
		assert.False(t, pile.CanStack(cards.NewCard(2, &suits.Diamonds)))
	})
}
