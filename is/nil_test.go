package is_test

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
)

func TestNil(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)
		matcher := assert.That(nil, is.Nil())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("someString", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Nil()
		assert.That(matcher.Matches("someString"), is.False())
		assert.That(matcher.Cause(), is.EqualTo("want <nil>; got=string(someString)"))
	})
}

func TestNotNil(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotNil()
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want not <nil>; got=<nil>"))
	})

	t.Run("someString", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That("someString", is.NotNil())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}
