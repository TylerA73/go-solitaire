package stack

import (
	"testing"

	"github.com/TylerA73/go-solitaire/cards"
	"github.com/TylerA73/go-solitaire/suits"
	"github.com/stretchr/testify/assert"
)

func Test_HomeCanStack(t *testing.T) {
	t.Run("should be able to stack card onto an empty home", func(t *testing.T) {
		home := NewHome(&suits.Clubs)

		assert.True(t, home.CanStack(card))
	})

	t.Run("should not be able to stack card onto an empty home", func(t *testing.T) {
		home := NewHome(&suits.Diamonds)

		assert.False(t, home.CanStack(card))
	})

	t.Run("should be able to stack card onto a home with cards", func(t *testing.T) {
		home := NewHome(&suits.Clubs)
		home.Add(*cards.NewCard(4, &suits.Clubs))

		assert.True(t, home.CanStack(cards.NewCard(5, &suits.Clubs)))
	})

	t.Run("should not be able to stack card onto a pile with cards", func(t *testing.T) {
		home := NewHome(&suits.Diamonds)
		home.Add(*cards.NewCard(4, &suits.Diamonds))

		assert.False(t, home.CanStack(cards.NewCard(5, &suits.Clubs)))
		assert.False(t, home.CanStack(cards.NewCard(4, &suits.Diamonds)))
		assert.False(t, home.CanStack(cards.NewCard(3, &suits.Diamonds)))
	})
}
