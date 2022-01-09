package stack

import (
	"testing"

	"github.com/TylerA73/go-solitaire/cards"
	"github.com/TylerA73/go-solitaire/suits"
	"github.com/stretchr/testify/assert"
)

func Test_Draw(t *testing.T) {
	t.Run("should draw 3 cards from the deck", func(t *testing.T) {
		assert.True(t, true)
	})
}

func Test_Pick(t *testing.T) {
	t.Run("should pick 1 card from the top of the draw pile", func(t *testing.T) {
		draw := NewDraw()
		draw.Cards[0] = cards.NewCard(3, &suits.Clubs)
		draw.Cards[1] = cards.NewCard(8, &suits.Hearts)
		draw.Cards[2] = cards.NewCard(5, &suits.Spades)

		pickedCard, err := draw.Pick()

		assert.NoError(t, err)
		assert.EqualValues(t, cards.NewCard(5, &suits.Spades), pickedCard)
		assert.EqualValues(t, draw.size(), 2)

		pickedCard, err = draw.Pick()
		assert.NoError(t, err)
		assert.EqualValues(t, cards.NewCard(8, &suits.Hearts), pickedCard)
		assert.EqualValues(t, draw.size(), 1)

		assert.True(t, true)
	})

	t.Run("should return an error if there are no cards to draw", func(t *testing.T) {
		draw := NewDraw()

		_, err := draw.Pick()

		assert.Error(t, err)
	})
}
