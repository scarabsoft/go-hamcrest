package is_test

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
)

func TestTrue(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.True()
		matched := matcher.Matches(true)

		assert.That(matched, is.True())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("false", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.True()
		matched := matcher.Matches(false)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want true; got false"))
	})
}

func TestFalse(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.False()
		matched := matcher.Matches(true)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want false; got true"))
	})

	t.Run("false", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.False()
		matched := matcher.Matches(false)

		assert.That(matched, is.True())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestOk(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Ok()
		matched := matcher.Matches(true)

		assert.That(matched, is.True())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("false", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Ok()
		matched := matcher.Matches(false)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want true; got false"))
	})
}

func TestNotOk(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotOk()
		matched := matcher.Matches(true)

		assert.That(matched, is.NotOk())
		assert.That(matcher.Cause(), is.EqualTo("want false; got true"))
	})

	t.Run("notOk", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotOk()
		matched := matcher.Matches(false)

		assert.That(matched, is.True())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}
