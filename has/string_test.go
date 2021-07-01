package has_test

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/has"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
)

func TestPrefix(t *testing.T) {

	t.Run("actual is nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Prefix("prefix")
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("actual not a string", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Prefix("prefix")
		matched := matcher.Matches(true)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [string]"))
	})

	t.Run("abcdef_bcd", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Prefix("bcd")
		matched := matcher.Matches("abcdef")

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want bcd to be prefix of abcdef; but was not"))
	})

	t.Run("abcdef_abcd", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That("abcde", has.Prefix("abcd"))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestNotPrefix(t *testing.T) {
	t.Run("actual is nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotPrefix("prefix")
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("actual not a string", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotPrefix("prefix")
		matched := matcher.Matches(true)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [string]"))
	})

	t.Run("abcdef_bcd", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That("abcde", has.NotPrefix("bcd"))
		assert.That(matcher.Cause(), is.EqualTo(""))

	})

	t.Run("abcdef_abcd", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotPrefix("abcd")
		matched := matcher.Matches("abcdef")

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want abcd not to be prefix of abcdef; but was"))
	})
}

func TestSuffix(t *testing.T) {
	t.Run("actual is nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Suffix("suffix")
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("actual not a string", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Suffix("suffix")
		matched := matcher.Matches(true)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [string]"))
	})

	t.Run("abcdef_bcd", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Suffix("bcd")
		matched := matcher.Matches("abcdef")

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want bcd to be suffix of abcdef; but was not"))
	})

	t.Run("abcdef_def", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That("abcdef", has.Suffix("def"))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestNotSuffix(t *testing.T) {
	t.Run("actual is nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotSuffix("suffix")
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("actual not a string", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotSuffix("suffix")
		matched := matcher.Matches(true)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [string]"))
	})

	t.Run("abcdef_bcd", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That("abcdef", has.NotSuffix("bcd"))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("abcdef_def", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotSuffix("def")
		matched := matcher.Matches("abcdef")

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want def not to be suffix of abcdef; but was"))
	})
}
