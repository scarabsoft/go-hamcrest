package is_test

import (
	"github.com/scarabsoft/go-hamcrest"
	"github.com/scarabsoft/go-hamcrest/is"
	"testing"
)

func TestEqualTo(t *testing.T) {

	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(nil, is.EqualTo(nil))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.EqualTo(nil)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value equal to int(1); got <nil>"))
	})

	t.Run("nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.EqualTo(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value equal to <nil>; got int(1)"))
	})

	t.Run("float32_float64", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.EqualTo(float32(13))
		matched := matcher.Matches(42.24)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("comparing different types; float64 != float32"))
	})

	t.Run("1_2", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.EqualTo(1)
		matched := matcher.Matches(2)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value equal to int(2); got int(1)"))
	})

	t.Run("42_42", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(42, is.EqualTo(42))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestNotEqual(t *testing.T) {

	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotEqualTo(nil)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value not equal to <nil>; got <nil>"))
	})

	t.Run("1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(1, is.NotEqualTo(nil))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(nil, is.NotEqualTo(1))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("float32_float64", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotEqualTo(float32(1))
		matched := matcher.Matches(42.24)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("comparing different types; float64 != float32"))
	})

	t.Run("1_2", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(1, is.NotEqualTo(2))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("42_42", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotEqualTo(42)
		matched := matcher.Matches(42)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value not equal to int(42); got int(42)"))
	})
}

func TestGreaterThan(t *testing.T) {

	t.Run("actual_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThan(1)
		matched := matcher.Matches([1]int{1})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [bool,string,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("given_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThan([1]int{1})
		matched := matcher.Matches(1)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given not one of [bool,string,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThan(nil)
		matched := matcher.Matches(nil)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThan(nil)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given is <nil>"))
	})

	t.Run("nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThan(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("int_float", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThan(1)
		matched := matcher.Matches(42.24)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("comparing different types; float64 != int"))
	})

	t.Run("1_2", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThan(2)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value greater than int(2); got int(1)"))
	})

	t.Run("2_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(2, is.GreaterThan(1))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("42_42", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThan(42)
		matched := matcher.Matches(42)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value greater than int(42); got int(42)"))
	})
}

func TestGreaterThanEqual(t *testing.T) {
	t.Run("actual_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThan(1)
		matched := matcher.Matches([1]int{1})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [bool,string,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("given_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThan([1]int{1})
		matched := matcher.Matches(1)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given not one of [bool,string,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThanEqual(nil)
		matched := matcher.Matches(nil)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThanEqual(nil)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given is <nil>"))
	})

	t.Run("nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThanEqual(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("int_float", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThanEqual(1)
		matched := matcher.Matches(42.24)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("comparing different types; float64 != int"))
	})

	t.Run("1_2", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.GreaterThanEqual(2)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value greater than equal to int(2); got int(1)"))
	})

	t.Run("2_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(2, is.GreaterThanEqual(1))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("42_42", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(42, is.GreaterThanEqual(42))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestLessThan(t *testing.T) {

	t.Run("actual_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThan(1)
		matched := matcher.Matches([1]int{1})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [bool,string,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("given_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThan([1]int{1})
		matched := matcher.Matches(1)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given not one of [bool,string,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThan(nil)
		matched := matcher.Matches(nil)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThan(nil)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given is <nil>"))
	})

	t.Run("nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThan(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("int_float", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThan(1)
		matched := matcher.Matches(42.24)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("comparing different types; float64 != int"))
	})

	t.Run("1_2", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(1, is.LessThan(2))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("2_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThan(1)
		matched := matcher.Matches(2)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value less than int(1); got int(2)"))
	})

	t.Run("42_42", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThan(42)
		matched := matcher.Matches(42)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value less than int(42); got int(42)"))
	})
}

func TestLessThanEqual(t *testing.T) {

	t.Run("actual_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThanEqual(1)
		matched := matcher.Matches([1]int{1})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [bool,string,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("given_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThanEqual([1]int{1})
		matched := matcher.Matches(1)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given not one of [bool,string,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThanEqual(nil)
		matched := matcher.Matches(nil)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThanEqual(nil)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given is <nil>"))
	})

	t.Run("nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThanEqual(1)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("int_float", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThanEqual(1)
		matched := matcher.Matches(42.24)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("comparing different types; float64 != int"))
	})

	t.Run("1_2", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(1, is.LessThanEqual(2))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("2_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThanEqual(1)
		matched := matcher.Matches(2)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value less than equal to int(1); got int(2)"))
	})

	t.Run("42_42", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.LessThanEqual(42)
		matched := matcher.Matches(42)

		assert.That(matched, is.True())
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestBetween(t *testing.T) {
	t.Run("actual_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Between(1, 2)
		matched := matcher.Matches([1]int{1})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("givenMin_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Between([1]int{1}, 2)
		matched := matcher.Matches(1)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMin not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("givenMax_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Between(1, [1]int{1})
		matched := matcher.Matches(1)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMax not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("nil_nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Between(nil, nil)
		matched := matcher.Matches(nil)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("nil_1_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Between(1, 2)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("1_nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Between(nil, 1)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMin is <nil>"))
	})

	t.Run("1_1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Between(1, nil)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMax is <nil>"))
	})

	t.Run("int_int_float", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Between(1, 1)
		matched := matcher.Matches(42.24)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("comparing different types; float64 != int"))
	})

	t.Run("1_2_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Between(2, 4)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value between int(2) and int(4); got int(1)"))
	})

	t.Run("1_4_2", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Between(4, 2)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMin int(4) must be <= givenMax int(2)"))
	})

	t.Run("2_1_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(2, is.Between(1, 4))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("2_2_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Between(2, 4)
		matched := matcher.Matches(2)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value between int(2) and int(4); got int(2)"))
	})

	t.Run("4_2_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.Between(2, 4)
		matched := matcher.Matches(4)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value between int(2) and int(4); got int(4)"))
	})
}

func TestBetweenOrEqual(t *testing.T) {
	t.Run("actual_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BetweenOrEqual(1, 2)
		matched := matcher.Matches([1]int{1})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("givenMin_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BetweenOrEqual([1]int{1}, 2)
		matched := matcher.Matches(1)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMin not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("givenMax_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BetweenOrEqual(1, [1]int{1})
		matched := matcher.Matches(1)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMax not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("nil_nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BetweenOrEqual(nil, nil)
		matched := matcher.Matches(nil)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("nil_1_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BetweenOrEqual(1, 2)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("1_nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BetweenOrEqual(nil, 1)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMin is <nil>"))
	})

	t.Run("1_1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BetweenOrEqual(1, nil)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMax is <nil>"))
	})

	t.Run("int_int_float", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BetweenOrEqual(1, 1)
		matched := matcher.Matches(42.24)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("comparing different types; float64 != int"))
	})

	t.Run("1_2_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BetweenOrEqual(2, 4)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value between or equal to int(2) and int(4); got int(1)"))
	})

	t.Run("1_4_2", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.BetweenOrEqual(4, 2)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMin int(4) must be <= givenMax int(2)"))
	})

	t.Run("2_1_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(2, is.BetweenOrEqual(1, 4))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("2_2_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(2, is.BetweenOrEqual(2, 4))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("4_2_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(4, is.BetweenOrEqual(2, 4))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

}

func TestNotBetween(t *testing.T) {
	t.Run("actual_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetween(1, 2)
		matched := matcher.Matches([1]int{1})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("givenMin_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetween([1]int{1}, 2)
		matched := matcher.Matches(1)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMin not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("givenMax_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetween(1, [1]int{1})
		matched := matcher.Matches(1)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMax not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("nil_nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetween(nil, nil)
		matched := matcher.Matches(nil)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("nil_1_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetween(1, 2)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("1_nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetween(nil, 1)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMin is <nil>"))
	})

	t.Run("1_1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetween(1, nil)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMax is <nil>"))
	})

	t.Run("int_int_float", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetween(1, 1)
		matched := matcher.Matches(42.24)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("comparing different types; float64 != int"))
	})

	t.Run("1_2_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(1, is.NotBetween(2, 4))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("1_4_2", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetween(4, 2)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMin int(4) must be <= givenMax int(2)"))
	})

	t.Run("2_1_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetween(1, 4)
		matched := matcher.Matches(2)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value not between int(1) and int(4); got int(2)"))
	})

	t.Run("2_2_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(2, is.NotBetween(2, 4))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("4_2_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(4, is.NotBetween(2, 4))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}

func TestNotBetweenOrEqualOrEqual(t *testing.T) {
	t.Run("actual_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetweenOrEqual(1, 2)
		matched := matcher.Matches([1]int{1})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("givenMin_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetweenOrEqual([1]int{1}, 2)
		matched := matcher.Matches(1)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMin not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("givenMax_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetweenOrEqual(1, [1]int{1})
		matched := matcher.Matches(1)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMax not one of [int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64]"))
	})

	t.Run("nil_nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetweenOrEqual(nil, nil)
		matched := matcher.Matches(nil)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("nil_1_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetweenOrEqual(1, 2)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("1_nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetweenOrEqual(nil, 1)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMin is <nil>"))
	})

	t.Run("1_1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetweenOrEqual(1, nil)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMax is <nil>"))
	})

	t.Run("int_int_float", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetweenOrEqual(1, 1)
		matched := matcher.Matches(42.24)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("comparing different types; float64 != int"))
	})

	t.Run("1_2_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(1, is.NotBetweenOrEqual(2, 4))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("1_4_2", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetweenOrEqual(4, 2)
		matched := matcher.Matches(1)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("givenMin int(4) must be <= givenMax int(2)"))
	})

	t.Run("2_1_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetweenOrEqual(1, 4)
		matched := matcher.Matches(2)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value not between or equal to int(1) and int(4); got int(2)"))
	})

	t.Run("2_2_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetweenOrEqual(2, 4)
		matched := matcher.Matches(2)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value not between or equal to int(2) and int(4); got int(2)"))
	})

	t.Run("4_2_4", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotBetweenOrEqual(2, 4)
		matched := matcher.Matches(4)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value not between or equal to int(2) and int(4); got int(4)"))
	})
}

func TestCloseTo(t *testing.T) {
	t.Run("actual_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.CloseTo(1.0, 2.0)
		matched := matcher.Matches([1]int{1})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [float32,float64]"))
	})

	t.Run("given_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.CloseTo("", 2.0)
		matched := matcher.Matches(1.0)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given not one of [float32,float64]"))
	})

	t.Run("error_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.CloseTo(1.0, "")
		matched := matcher.Matches(1.0)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("error not one of [float32,float64]"))
	})

	t.Run("nil_nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.CloseTo(nil, nil)
		matched := matcher.Matches(nil)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("nil_1_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.CloseTo(1.0, 2.0)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("1_nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.CloseTo(nil, 1.0)
		matched := matcher.Matches(1.0)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given is <nil>"))
	})

	t.Run("1_1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.CloseTo(1.0, nil)
		matched := matcher.Matches(1.0)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("error is <nil>"))
	})

	t.Run("float64_float64_float32", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.CloseTo(1.0, float32(1))
		matched := matcher.Matches(42.24)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("comparing different types; float64 != float32"))
	})

	t.Run("4.6_4_0.5", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.CloseTo(4.0, 0.5)
		matched := matcher.Matches(4.6)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value close to float64(4) with error float64(0.5); got float64(4.6)"))
	})

	t.Run("4.5_4_0.5", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(4.5, is.CloseTo(4.0, 0.5))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("3.5_4_0.5", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(3.5, is.CloseTo(4.0, 0.5))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("3.4_4_0.5", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.CloseTo(4.0, 0.5)
		matched := matcher.Matches(3.4)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value close to float64(4) with error float64(0.5); got float64(3.4)"))
	})
}

func TestNotCloseTo(t *testing.T) {
	t.Run("actual_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotCloseTo(1.0, 2.0)
		matched := matcher.Matches([1]int{1})
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual not one of [float32,float64]"))
	})

	t.Run("given_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotCloseTo("", 2.0)
		matched := matcher.Matches(1.0)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given not one of [float32,float64]"))
	})

	t.Run("error_unsupported", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotCloseTo(1.0, "")
		matched := matcher.Matches(1.0)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("error not one of [float32,float64]"))
	})

	t.Run("nil_nil_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotCloseTo(nil, nil)
		matched := matcher.Matches(nil)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("nil_1_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotCloseTo(1.0, 2.0)
		matched := matcher.Matches(nil)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("actual is <nil>"))
	})

	t.Run("1_nil_1", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotCloseTo(nil, 1.0)
		matched := matcher.Matches(1.0)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("given is <nil>"))
	})

	t.Run("1_1_nil", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotCloseTo(1.0, nil)
		matched := matcher.Matches(1.0)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("error is <nil>"))
	})

	t.Run("float64_float64_float32", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotCloseTo(1.0, float32(1))
		matched := matcher.Matches(42.24)

		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("comparing different types; float64 != float32"))
	})

	t.Run("4.6_4_0.5", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(4.6, is.NotCloseTo(4.0, 0.5))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})

	t.Run("4.5_4_0.5", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotCloseTo(4.0, 0.5)
		matched := matcher.Matches(4.5)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value not close to float64(4) with error float64(0.5); got float64(4.5)"))
	})

	t.Run("3.5_4_0.5", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := is.NotCloseTo(4.0, 0.5)
		matched := matcher.Matches(3.5)
		assert.That(matched, is.False())
		assert.That(matcher.Cause(), is.EqualTo("want value not close to float64(4) with error float64(0.5); got float64(3.5)"))
	})

	t.Run("3.4_4_0.5", func(t *testing.T) {
		assert := hamcrest.NewAssertion(t)

		matcher := assert.That(3.4, is.NotCloseTo(4.0, 0.5))
		assert.That(matcher.Cause(), is.EqualTo(""))
	})
}
