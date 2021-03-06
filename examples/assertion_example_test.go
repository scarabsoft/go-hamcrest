package examples

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
)

func TestAssertions(t *testing.T) {
	t.Run("assert that nil is nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		assert.That(nil, is.Nil())
	})

	t.Run("assert that true is true", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		assert.That(true, is.True())
	})

	t.Run("assert that false is false", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		assert.That(false, is.False())
	})
}
