package is_test

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
	"time"
)

func createTime(v string) time.Time {
	result, err := time.Parse(time.RFC3339, v)
	if err != nil {
		panic(err)
	}
	return result
}

var (
	someTime       = createTime("2012-11-02T22:08:41+00:00")
	beforeThatTime = createTime("2012-11-01T22:08:41+00:00")
	afterThatTime  = createTime("2012-11-03T22:08:41+00:00")
)

func TestBefore(t *testing.T) {
	t.Run("nil_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Before(someTime)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("someTime_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Before(nil)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("expected is <nil>"))
	})

	t.Run("int_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Before(someTime)
		matched := matcher.Matches(10)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want actual to be type of time.Time; got int(10)"))
	})

	t.Run("someTime_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Before(10)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want expected to be type of time.Time; got int(10)"))
	})

	t.Run("beforeThatTime_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(beforeThatTime, is.Before(someTime))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("someTime_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Before(someTime)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value before 2012-11-02T22:08:41Z; got 2012-11-02T22:08:41Z"))
	})
}

func TestNotBefore(t *testing.T) {
	t.Run("nil_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBefore(someTime)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("someTime_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBefore(nil)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("expected is <nil>"))
	})

	t.Run("int_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBefore(someTime)
		matched := matcher.Matches(10)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want actual to be type of time.Time; got int(10)"))
	})

	t.Run("someTime_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBefore(10)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want expected to be type of time.Time; got int(10)"))
	})

	t.Run("beforeThatTime_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBefore(someTime)
		matched := matcher.Matches(beforeThatTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value not before 2012-11-02T22:08:41Z; got 2012-11-01T22:08:41Z"))
	})

	t.Run("someTime_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(someTime, is.NotBefore(someTime))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestBeforeOrEqual(t *testing.T) {
	t.Run("nil_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BeforeOrEqual(someTime)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("someTime_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BeforeOrEqual(nil)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("expected is <nil>"))
	})

	t.Run("int_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BeforeOrEqual(someTime)
		matched := matcher.Matches(10)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want actual to be type of time.Time; got int(10)"))
	})

	t.Run("someTime_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BeforeOrEqual(10)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want expected to be type of time.Time; got int(10)"))
	})

	t.Run("BeforeOrSameThatTime_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(beforeThatTime, is.BeforeOrEqual(someTime))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("someTime_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(beforeThatTime, is.BeforeOrEqual(someTime))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestAfter(t *testing.T) {
	t.Run("nil_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.After(someTime)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("someTime_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.After(nil)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("expected is <nil>"))
	})

	t.Run("int_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.After(someTime)
		matched := matcher.Matches(10)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want actual to be type of time.Time; got int(10)"))
	})

	t.Run("someTime_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.After(10)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want expected to be type of time.Time; got int(10)"))
	})

	t.Run("afterThatTime_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(afterThatTime, is.After(someTime))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("someTime_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.After(someTime)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value after 2012-11-02T22:08:41Z; got 2012-11-02T22:08:41Z"))
	})
}

func TestNotAfter(t *testing.T) {
	t.Run("nil_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotAfter(someTime)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("someTime_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotAfter(nil)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("expected is <nil>"))
	})

	t.Run("int_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotAfter(someTime)
		matched := matcher.Matches(10)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want actual to be type of time.Time; got int(10)"))
	})

	t.Run("someTime_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotAfter(10)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want expected to be type of time.Time; got int(10)"))
	})

	t.Run("afterThatTime_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotAfter(someTime)
		matched := matcher.Matches(afterThatTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value not after 2012-11-02T22:08:41Z; got 2012-11-03T22:08:41Z"))
	})

	t.Run("someTime_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(someTime, is.NotAfter(someTime))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestAfterOrEqual(t *testing.T) {
	t.Run("nil_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.AfterOrEqual(someTime)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("someTime_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.AfterOrEqual(nil)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("expected is <nil>"))
	})

	t.Run("int_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.AfterOrEqual(someTime)
		matched := matcher.Matches(10)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want actual to be type of time.Time; got int(10)"))
	})

	t.Run("someTime_int", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.AfterOrEqual(10)
		matched := matcher.Matches(someTime)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want expected to be type of time.Time; got int(10)"))
	})

	t.Run("afterThatTime_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(afterThatTime, is.AfterOrEqual(someTime))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("someTime_someTime", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(someTime, is.AfterOrEqual(someTime))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}
