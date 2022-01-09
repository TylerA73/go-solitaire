package cards

import (
	"testing"

	"github.com/TylerA73/go-solitaire/suits"
	"github.com/stretchr/testify/assert"
)

var (
	c1 *Card = NewCard(5, &suits.Hearts)
)

func Test_Card(t *testing.T) {
	t.Run("should return true if card suit is red", func(t *testing.T) {
		assert.True(t, c1.IsRed())
	})

	t.Run("should return false if the suit is not red", func(t *testing.T) {
		c2 := NewCard(5, &suits.Clubs)

		assert.False(t, c2.IsRed())
	})
}
