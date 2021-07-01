package is_test

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
)

func TestEmpty(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Empty()
		assert.False(matcher.Matches(nil))
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("[3]int{1,2,3}", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Empty()
		assert.False(matcher.Matches([3]int{1, 2, 3}))
		assert.That(matcher.Cause(), is.EqualTo("want to be empty; got length 3"))
	})

	t.Run("empty_string", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That("", is.Empty())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("someString", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Empty()
		assert.False(matcher.Matches("someString"))
		assert.That(matcher.Cause(), is.EqualTo("want to be empty; got length 10"))
	})

	t.Run("[]int{}", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([]int{}, is.Empty())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[]int{1,2,3}", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Empty()
		assert.False(matcher.Matches([]int{1, 2, 3}))
		assert.That(matcher.Cause(), is.EqualTo("want to be empty; got length 3"))
	})

	t.Run("empty chan", func(t *testing.T) {
		c := make(chan int, 2)

		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(c, is.Empty())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("not empty chan", func(t *testing.T) {
		c := make(chan int, 2)
		c <- 42

		assert := hamcrest.NewAssertion(t)

		matcher := is.Empty()
		assert.False(matcher.Matches(c))
		assert.That(matcher.Cause(), is.EqualTo("want to be empty; got length 1"))
	})
}

func TestNotEmpty(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotEmpty()
		assert.False(matcher.Matches(nil))
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("[3]int{1,2,3}", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([3]int{1, 2, 3}, is.NotEmpty())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("empty_string", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotEmpty()
		assert.False(matcher.Matches(""))
		assert.That(matcher.Cause(), is.EqualTo("want not to be empty; got empty"))
	})

	t.Run("someString", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That("someString", is.NotEmpty())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[]int{}", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotEmpty()
		assert.False(matcher.Matches([]int{}))
		assert.That(matcher.Cause(), is.EqualTo("want not to be empty; got empty"))
	})

	t.Run("[]int{1,2,3}", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([]int{1, 2, 3}, is.NotEmpty())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("empty chan", func(t *testing.T) {
		c := make(chan int, 2)

		assert := hamcrest.NewAssertion(t)

		matcher := is.NotEmpty()
		assert.False(matcher.Matches(c))
		assert.That(matcher.Cause(), is.EqualTo("want not to be empty; got empty"))
	})

	t.Run("not empty chan", func(t *testing.T) {
		c := make(chan int, 2)
		c <- 42

		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(c, is.NotEmpty())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}
