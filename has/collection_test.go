package has_test

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/has"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
)

func TestLength(t *testing.T) {
	t.Run("length_is_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Length(nil)
		matched := matcher.Matches([]string{"given"})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("length is <nil>"))
	})

	t.Run("length_is_not_integer", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Length(32.1)
		matched := matcher.Matches([]string{"given"})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("length not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64]"))
	})

	t.Run("actual_is_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Length(3)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("actual_is_unsupported_type", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Length(3)
		matched := matcher.Matches(struct{}{})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [array,chan,map,string,slice,ptr]"))
	})

	t.Run("length_match_unsigned_integer_length", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([1]string{"given"}, has.Length(uint16(1)))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("actual_match", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([1]string{"given"}, has.Length(1))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("actual_no_match", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Length(2)
		matched := matcher.Matches([1]string{"given"})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want length of 2; got 1"))
	})

}

func TestItem(t *testing.T) {
	t.Run("actual_is_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Item(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("[1,2,3]_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([]int{1, 2, 3}, has.Item(1))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("&[1,2,3]_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(&[]int{1, 2, 3}, has.Item(1))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[abcd]_abcd", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([]string{"abcd"}, has.Item("abcd"))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[1,2,3]_true", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Item(true)
		matched := matcher.Matches([]int{1, 2, 3})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want bool(true) to be part of []int([1 2 3]); but was not"))
	})

	t.Run("[1,2,3]_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Item(nil)
		matched := matcher.Matches([]int{1, 2, 3})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want <nil> to be part of []int([1 2 3]); but was not"))
	})

	t.Run("[nil,*someString]_nil", func(t *testing.T) {
		var someString = "someString"
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([]*string{nil, &someString}, has.Item(nil))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[nil,*someString]_*someString", func(t *testing.T) {
		var someString = "someString"
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([]*string{nil, &someString}, has.Item(&someString))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestNotItem(t *testing.T) {

	t.Run("actual_is_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotItem(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("[1,2,3]_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotItem(1)
		matched := matcher.Matches([]int{1, 2, 3})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want int(1) to not be part of []int([1 2 3]); but was"))
	})

	t.Run("&[1,2,3]_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotItem(1)
		matched := matcher.Matches(&[]int{1, 2, 3})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want int(1) to not be part of *[]int(&[1 2 3]); but was"))
	})

	t.Run("[1,2,3]_true", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([]int{1, 2, 3}, has.NotItem(true))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[1,2,3]_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([]int{1, 2, 3}, has.NotItem(4))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestItems(t *testing.T) {
	t.Run("actual_is_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Items(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("[1,2,3]_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([]int{1, 2, 3}, has.Items(1))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[1,2,3]_[1,2]", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([]int{1, 2, 3}, has.Items(1))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[1,2,3]_[1,2,3,4]", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Items(1, 2, 3, 4)
		matched := matcher.Matches([]int{1, 2, 3})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want int(4) to be part of []int([1 2 3]); but was not"))
	})
}

func TestNotItems(t *testing.T) {
	t.Run("actual_is_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotItems(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("[1,2,3]_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotItems(1)
		matched := matcher.Matches([]int{1, 2, 3})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want int(1) to not be part of []int([1 2 3]); but was"))
	})

	t.Run("[1,2,3]_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That([]int{1, 2, 3}, has.NotItems(4))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[1,2]_[4,5,6,1]", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotItems(4, 5, 6, 1)
		matched := matcher.Matches([]int{1, 2})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want int(1) to not be part of []int([1 2]); but was"))
	})
}

func TestKey(t *testing.T) {

	t.Run("actual_is_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Key(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("actual_is_not_a_map", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Key(1)
		matched := matcher.Matches([]int{10, 20})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [map]"))
	})

	t.Run("[10:10,20:20]_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Key(1)
		matched := matcher.Matches(map[int]int{10: 10, 20: 20})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want int(1) to be key of map[int]int(map[10:10 20:20]); but was not"))
	})

	t.Run("[10:10,20:20]_10", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(map[int]int{10: 10, 20: 20}, has.Key(10))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[1:1]_true", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Key(true)
		matched := matcher.Matches(map[int]int{1: 1})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want bool(true) to be key of map[int]int(map[1:1]); but was not"))
	})
}

func TestNotKey(t *testing.T) {
	t.Run("actual_is_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotKey(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("actual_is_not_a_map", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotKey(1)
		matched := matcher.Matches([]int{10, 20})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [map]"))
	})

	t.Run("[10:10,20:20]_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(map[int]int{10: 10, 20: 20}, has.NotKey(1))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[10:10,20:20]_10", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotKey(10)
		matched := matcher.Matches(map[int]int{10: 10, 20: 20})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want int(10) not to be key of map[int]int(map[10:10 20:20]); but was"))
	})

	t.Run("[1:1]_true", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(map[int]int{1: 1}, has.NotKey(true))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestKeys(t *testing.T) {
	t.Run("actual_is_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Keys(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("actual_is_not_a_map", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Keys(1)
		matched := matcher.Matches([]int{10, 20})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [map]"))
	})

	t.Run("[10:10,20:20]_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Keys(1)
		matched := matcher.Matches(map[int]int{10: 10, 20: 20})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want int(1) to be key of map[int]int(map[10:10 20:20]); but was not"))
	})

	t.Run("[10:10,20:20]_10", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(map[int]int{10: 10, 20: 20}, has.Keys(10))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[10:10,20:20]_[10,20]", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(map[int]int{10: 10, 20: 20}, has.Keys(10, 20))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[1:1]_true", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.Keys(true)
		matched := matcher.Matches(map[int]int{1: 1})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want bool(true) to be key of map[int]int(map[1:1]); but was not"))
	})
}

func TestNotKeys(t *testing.T) {
	t.Run("actual_is_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotKeys(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("actual_is_not_a_map", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotKeys(1)
		matched := matcher.Matches([]int{10, 20})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [map]"))
	})

	t.Run("[10:10,20:20]_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(map[int]int{10: 10, 20: 20}, has.NotKeys(1))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("[10:10,20:20]_10", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotKeys(10)
		matched := matcher.Matches(map[int]int{10: 10, 20: 20})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want int(10) not to be key of map[int]int(map[10:10 20:20]); but was"))
	})

	t.Run("[10:10,20:20]_[1,2,3,20]", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := has.NotKeys(1, 2, 3, 20)
		matched := matcher.Matches(map[int]int{10: 10, 20: 20})

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want int(20) not to be key of map[int]int(map[10:10 20:20]); but was"))
	})

	t.Run("[1:1]_true", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(map[int]int{1: 1}, has.NotKeys(true))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}
