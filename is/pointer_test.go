package is_test

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
)

func TestPointingTo(t *testing.T) {
	var expectedPtr = new(int)
	*expectedPtr = 10

	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(nil, is.PointingTo(nil))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("ptr_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.PointingTo(nil)
		matched := matcher.Matches(expectedPtr)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("expected is <nil>"))
	})

	t.Run("nil_ptr", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.PointingTo(expectedPtr)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("1_ptr", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.PointingTo(expectedPtr)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [ptr,unsafe.Pointer]"))
	})

	t.Run("ptr_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.PointingTo(1)
		matched := matcher.Matches(expectedPtr)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("expected not one of [ptr,unsafe.Pointer]"))
	})

	t.Run("same_ptr", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(expectedPtr, is.PointingTo(expectedPtr))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("different_ptr_variables_same_address", func(t *testing.T) {
		anotherPtr := expectedPtr

		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(anotherPtr, is.PointingTo(expectedPtr))
		assert.That(matcher.Cause(), is.EqualTo(""))

		matcher = assert.That(expectedPtr, is.PointingTo(anotherPtr))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("same_ptr_type_different_addresses", func(t *testing.T) {
		var anotherPtr = new(int)
		*anotherPtr = 10

		assert := hamcrest.NewAssertion(t)

		matcher := is.PointingTo(expectedPtr)
		matched := matcher.Matches(anotherPtr)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.MatchingPattern(`to\s\*int\([a-zA-z0-9]*\);\sgot\s\*int\([a-zA-z0-9]*\)`))
	})

	t.Run("different_types", func(t *testing.T) {
		var anotherPtr = new(float32)
		*anotherPtr = float32(42)

		assert := hamcrest.NewAssertion(t)

		matcher := is.PointingTo(expectedPtr)
		matched := matcher.Matches(anotherPtr)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.MatchingPattern(`to\s\*float32\([a-zA-z0-9]*\);\sgot\s\*int\([a-zA-z0-9]*\)`))
	})
}

func TestNotPointingTo(t *testing.T) {
	var expectedPtr = new(int)
	*expectedPtr = 10

	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(nil, is.NotPointingTo(nil))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("ptr_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotPointingTo(nil)
		matched := matcher.Matches(expectedPtr)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("expected is <nil>"))
	})

	t.Run("nil_ptr", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotPointingTo(expectedPtr)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("1_ptr", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotPointingTo(expectedPtr)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [ptr,unsafe.Pointer]"))
	})

	t.Run("ptr_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotPointingTo(1)
		matched := matcher.Matches(expectedPtr)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("expected not one of [ptr,unsafe.Pointer]"))
	})

	t.Run("same_ptr", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotPointingTo(expectedPtr)
		matched := matcher.Matches(expectedPtr)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.MatchingPattern(`want\spointer\sto\snot\spoint\sto\s\*int\([a-zA-z0-9]*\);\sgot\s\*int\([a-zA-z0-9]*\)`))
	})

	t.Run("different_ptr_variables_same_address", func(t *testing.T) {
		anotherPtr := expectedPtr

		assert := hamcrest.NewAssertion(t)

		matcher := is.NotPointingTo(expectedPtr)
		matched := matcher.Matches(anotherPtr)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.MatchingPattern(`want\spointer\sto\snot\spoint\sto\s\*int\([a-zA-z0-9]*\);\sgot\s\*int\([a-zA-z0-9]*\)`))

		matcher = is.NotPointingTo(anotherPtr)
		matched = matcher.Matches(expectedPtr)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.MatchingPattern(`want\spointer\sto\snot\spoint\sto\s\*int\([a-zA-z0-9]*\);\sgot\s\*int\([a-zA-z0-9]*\)`))
	})

	t.Run("same_ptr_type_different_addresses", func(t *testing.T) {
		var anotherPtr = new(int)
		*anotherPtr = 10

		assert := hamcrest.NewAssertion(t)

		matched := assert.That(anotherPtr, is.NotPointingTo(expectedPtr))
		assert.That(matched.Cause(), is.EqualTo(""))
	})

	t.Run("different_types", func(t *testing.T) {
		var anotherPtr = new(float32)
		*anotherPtr = float32(42)

		assert := hamcrest.NewAssertion(t)

		matched := assert.That(anotherPtr, is.NotPointingTo(expectedPtr))
		assert.That(matched.Cause(), is.EqualTo(""))
	})
}
