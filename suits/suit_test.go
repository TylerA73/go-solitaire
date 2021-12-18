package suits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsEqual(t *testing.T) {
	var (
		s1 *Suit = &Hearts
		s2 *Suit
	)

	t.Run("should be equal if the suits are the same", func(t *testing.T) {
		s2 = &Hearts

		assert.True(t, s1.IsEqual(s2))
	})

	t.Run("should not be equal if the suits are not the same, but are the same colour", func(t *testing.T) {
		s2 = &Diamonds

		assert.False(t, s1.IsEqual(s2))
	})

	t.Run("should not be equal if the suits are not the same, and are not the same colour", func(t *testing.T) {
		s2 = &Spades

		assert.False(t, s1.IsEqual(s2))
	})
}
