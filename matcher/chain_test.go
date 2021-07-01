package matcher_test

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/is"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"testing"
)

func TestMatchIfAllAreNil(t *testing.T) {

	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.MatchIfAllAreNil(nil, nil)()

		assert.That(result.Matched, is.True())
		assert.That(result.Cause, is.EqualTo(""))
	})

	t.Run("1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.MatchIfAllAreNil(1, nil)()

		assert.That(result.Matched, is.True())
		assert.That(result.Cause, is.EqualTo(""))
	})

	t.Run("nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.MatchIfAllAreNil(nil, 1)()

		assert.That(result.Matched, is.True())
		assert.That(result.Cause, is.EqualTo(""))
	})

	t.Run("1_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.MatchIfAllAreNil(1, 1)()

		assert.That(result.Matched, is.True())
		assert.That(result.Cause, is.EqualTo(""))
	})
}

func TestFailIfIsNil(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfIsNil("someName", nil)()

		assert.That(result.Matched, is.False())
		assert.That(result.Cause, is.EqualTo("someName is <nil>"))
	})

	t.Run("1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfIsNil("someName", 1)()

		assert.That(result.Matched, is.True())
		assert.That(result.Cause, is.EqualTo(""))
	})
}

func TestFailIfAnyIsNil(t *testing.T) {
	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfAnyIsNil("some message %s %s", nil, nil)()

		assert.That(result.Matched, is.False())
		assert.That(result.Cause, is.EqualTo("some message <nil> <nil>"))
	})

	t.Run("1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfAnyIsNil("some message %s %s", 1, nil)()

		assert.That(result.Matched, is.False())
		assert.That(result.Cause, is.EqualTo("some message int(1) <nil>"))
	})

	t.Run("nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfAnyIsNil("some message %s %s", nil, 1)()

		assert.That(result.Matched, is.False())
		assert.That(result.Cause, is.EqualTo("some message <nil> int(1)"))
	})

	t.Run("1_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfAnyIsNil("some message %s %s", 1, 1)()

		assert.That(result.Matched, is.True())
		assert.That(result.Cause, is.EqualTo(""))
	})
}

func TestFailIfRestrictedType(t *testing.T) {
	t.Run("allowed", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfRestrictedType("someName", 1, internal.RestrictedToNumericStringBoolKinds)()

		assert.That(result.Matched, is.True())
		assert.That(result.Cause, is.EqualTo(""))
	})

	t.Run("not_allowed", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfRestrictedType("someName", "someValue", internal.RestrictedToBoolKind)()
		assert.That(result.Matched, is.False())
		assert.That(result.Cause, is.EqualTo("someName not one of [bool]"))
	})
}

func TestFailIfNotSameType(t *testing.T) {
	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfNotSameType(nil, nil)()

		assert.That(result.Matched, is.True())
		assert.That(result.Cause, is.EqualTo(""))
	})

	t.Run("nil_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfNotSameType(nil, 1)()

		assert.That(result.Matched, is.False())
		assert.That(result.Cause, is.EqualTo("comparing different types; <nil> != int"))
	})

	t.Run("int_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfNotSameType(1, nil)()

		assert.That(result.Matched, is.False())
		assert.That(result.Cause, is.EqualTo("comparing different types; int != <nil>"))
	})

	t.Run("int_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfNotSameType(1, 1)()

		assert.That(result.Matched, is.True())
		assert.That(result.Cause, is.EqualTo(""))
	})

	t.Run("float_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		result := matcher.FailIfNotSameType(1, 42.42)()

		assert.That(result.Matched, is.False())
		assert.That(result.Cause, is.EqualTo("comparing different types; int != float64"))
	})
}

func TestChain(t *testing.T) {
	t.Run("empty_chain_matches", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		instance := matcher.New(func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain
		})

		assert.That(instance.Matches(10), is.True())
		assert.That(instance.Cause(), is.EqualTo(""))
	})

	t.Run("failed_match_interrupts_chain", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		firstStageCalled, secondStageCalled := false, false

		instance := matcher.New(func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(func() matcher.MatchResult {
					firstStageCalled = true
					return matcher.Failed("failed_stage")
				}).
				Add(func() matcher.MatchResult {
					secondStageCalled = true
					return matcher.Matched()
				})
		})

		assert.That(instance.Matches("abc"), is.False())
		assert.That(instance.Cause(), is.EqualTo("failed_stage"))

		assert.That(firstStageCalled, is.True())
		assert.That(secondStageCalled, is.False())
	})

	t.Run("successful_match_interrupts_chain", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		firstStageCalled, secondStageCalled := false, false

		instance := matcher.New(func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(func() matcher.MatchResult {
					firstStageCalled = true
					return matcher.Matched()
				}).
				Add(func() matcher.MatchResult {
					secondStageCalled = true
					return matcher.Failed("I_never_get_called")
				})
		})

		assert.That(instance.Matches("abc"), is.True())
		assert.That(instance.Cause(), is.EqualTo(""))

		assert.That(firstStageCalled, is.True())
		assert.That(secondStageCalled, is.False())
	})

	t.Run("chain_without_matches_or_failures_continues_and_finishes_successful", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		firstStageCalled, secondStageCalled := false, false

		instance := matcher.New(func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(func() matcher.MatchResult {
					firstStageCalled = true
					return matcher.Next()
				}).
				Add(func() matcher.MatchResult {
					secondStageCalled = true
					return matcher.Next()
				})
		})

		assert.That(instance.Matches("abc"), is.True())
		assert.That(instance.Cause(), is.EqualTo(""))

		assert.That(firstStageCalled, is.True())
		assert.That(secondStageCalled, is.True())
	})
}
