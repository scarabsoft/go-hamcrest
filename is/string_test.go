package is_test

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
)

const pattern = "\\w+\\s\\d+\\s\\w+"
const matchingString = "Test 1234 Pattern"
const notMatchingString = "Not Matching Pattern"

func TestMatchingPattern(t *testing.T) {
	t.Run("nil_pattern", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.MatchingPattern(pattern)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("matchingString_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.MatchingPattern(nil)
		matched := matcher.Matches(matchingString)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given is <nil>"))
	})

	t.Run("int_pattern", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.MatchingPattern(pattern)
		matched := matcher.Matches(42)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [string]"))
	})

	t.Run("matchingString_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.MatchingPattern(42)
		matched := matcher.Matches(matchingString)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given not one of [string]"))
	})

	t.Run("not_matching", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.MatchingPattern(pattern)
		matched := matcher.Matches(notMatchingString)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value to match \\w+\\s\\d+\\s\\w+; got Not Matching Pattern"))
	})

	t.Run("matching", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(matchingString, is.MatchingPattern(pattern))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestNotMatchingPattern(t *testing.T) {
	t.Run("nil_pattern", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotMatchingPattern(pattern)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("matchingString_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotMatchingPattern(nil)
		matched := matcher.Matches(matchingString)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given is <nil>"))
	})

	t.Run("int_pattern", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotMatchingPattern(pattern)
		matched := matcher.Matches(42)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [string]"))
	})

	t.Run("matchingString_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotMatchingPattern(42)
		matched := matcher.Matches(matchingString)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given not one of [string]"))
	})

	t.Run("not_matching", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(notMatchingString, is.NotMatchingPattern(pattern))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("matching", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotMatchingPattern(pattern)
		matched := matcher.Matches(matchingString)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value not to match \\w+\\s\\d+\\s\\w+; got Test 1234 Pattern"))
	})
}
