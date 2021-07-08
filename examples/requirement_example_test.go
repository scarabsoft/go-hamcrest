package examples

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
)

func TestRequirements(t *testing.T) {
	t.Run("require that nil is nil", func(t *testing.T) {
		require := hamcrest.NewRequirement(t)

		require.That(nil, is.Nil())
	})

	t.Run("require that true is true", func(t *testing.T) {
		require := hamcrest.NewRequirement(t)

		require.That(true, is.True())
	})

	t.Run("require that false is false", func(t *testing.T) {
		require := hamcrest.NewRequirement(t)

		require.That(false, is.False())
	})
}
