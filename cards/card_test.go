package cards

import (
	"testing"

	"github.com/TylerA73/go-solitaire/suits"
	"github.com/stretchr/testify/assert"
)

var (
	c1 *Card = NewCard(5, &suits.Hearts)
)

func Test(t *testing.T) {
	t.Run("should return true if card suit is red", func(t *testing.T) {
		assert.True(t, c1.IsRed())
	})

	t.Run("should return false if the suit is not red", func(t *testing.T) {
		c2 := NewCard(5, &suits.Clubs)

		assert.False(t, c2.IsRed())
	})

	t.Run("should return true if c2 and c3 can be stacked on top of c1 in a pile", func(t *testing.T) {
		c2 := NewCard(4, &suits.Clubs)
		c3 := NewCard(4, &suits.Spades)

		assert.True(t, c1.CanStackOnPile(c2))
		assert.True(t, c1.CanStackOnPile(c3))
	})

	t.Run("should return false if c2 and c3 can't be stacked on top of c1 in a pile", func(t *testing.T) {
		c2 := NewCard(8, &suits.Clubs)
		c3 := NewCard(8, &suits.Diamonds)

		assert.False(t, c1.CanStackOnPile(c2))
		assert.False(t, c1.CanStackOnPile(c3))
	})

	t.Run("should return true if c2 can be stacked on top of c1 in a home pile", func(t *testing.T) {
		c2 := NewCard(6, &suits.Hearts)

		assert.True(t, c1.CanStackOnHome(c2))
	})

	t.Run("should return false if c2 can't be stacked on top of c1 in a home pile", func(t *testing.T) {
		c2 := NewCard(6, &suits.Clubs)

		assert.False(t, c1.CanStackOnHome(c2))
	})
}
