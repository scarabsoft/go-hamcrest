package has_test

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/has"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
)

func TestSameTypeAs(t *testing.T) {
	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(nil, has.SameTypeAs(nil))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("int_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(42, has.SameTypeAs(1337))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("*int_*int", func(t *testing.T) {
		someInt := 42
		anotherInt := 1337

		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(&someInt, has.SameTypeAs(&anotherInt))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("int_float64", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.SameTypeAs(42.42)

		assert.False(matcher.Matches(1))
		assert.That(matcher.Cause(), is.EqualTo("want same type as int; got float64"))
	})

	t.Run("*int_*float64", func(t *testing.T) {
		someInt := 42
		someFloat := 13.37

		assert := hamcrest.NewAssertion(t)

		matcher := has.SameTypeAs(&someFloat)

		assert.False(matcher.Matches(&someInt))
		assert.That(matcher.Cause(), is.EqualTo("want same type as *int; got *float64"))
	})
}

func TestNotSameTypeAs(t *testing.T) {
	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotSameTypeAs(nil)
		assert.False(matcher.Matches(nil))
		assert.That(matcher.Cause(), is.EqualTo("want not same type as <nil>; got <nil>"))
	})

	t.Run("int_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotSameTypeAs(1337)
		assert.False(matcher.Matches(42))
		assert.That(matcher.Cause(), is.EqualTo("want not same type as int; got int"))
	})

	t.Run("*int_*int", func(t *testing.T) {
		someInt := 42
		anotherInt := 1337

		assert := hamcrest.NewAssertion(t)

		matcher := has.NotSameTypeAs(&anotherInt)
		assert.False(matcher.Matches(&someInt))
		assert.That(matcher.Cause(), is.EqualTo("want not same type as *int; got *int"))
	})

	t.Run("int_float64", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(1, has.NotSameTypeAs(42.42))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("*int_*float64", func(t *testing.T) {
		someInt := 42
		someFloat := 13.37

		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(&someInt, has.NotSameTypeAs(&someFloat))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}