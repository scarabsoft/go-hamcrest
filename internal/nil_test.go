package internal

import (
	"testing"
)

type testImpl struct{}
type test interface {
	test()
}

func (t *testImpl) test() {}

func TestIsNil(t *testing.T) {
	t.Run("nil_slice", func(t *testing.T) {
		var expectedNilSlice []string = nil
		matches := IsNil(expectedNilSlice)
		assertTrue(t, matches)
	})

	t.Run("not_nil_slice", func(t *testing.T) {
		var expectedNilSlice = make([]string, 0)
		matches := IsNil(expectedNilSlice)
		assertFalse(t, matches)
	})

	t.Run("nil_function", func(t *testing.T) {
		var expectedNilSlice func() = nil
		matches := IsNil(expectedNilSlice)
		assertTrue(t, matches)
	})

	t.Run("not_nil_function", func(t *testing.T) {
		var expectedNilSlice = func() {}
		matches := IsNil(expectedNilSlice)
		assertFalse(t, matches)
	})

	t.Run("nil_channel", func(t *testing.T) {
		f := func(x chan string) {
			matches := IsNil(x)
			assertTrue(t, matches)
		}
		f(nil)
	})

	t.Run("not_nil_channel", func(t *testing.T) {
		f := func(x chan string) {
			matches := IsNil(x)
			assertFalse(t, matches)
		}
		f(make(chan string))
	})

	t.Run("nil_interface", func(t *testing.T) {
		f := func(x interface{}) {
			matches := IsNil(x)
			assertTrue(t, matches)
		}
		f(nil)
	})

	t.Run("not_nil_interface", func(t *testing.T) {
		f := func(x interface{}) {
			matches := IsNil(x)
			assertFalse(t, matches)
		}
		f(&testImpl{})
	})

	t.Run("nil_map", func(t *testing.T) {
		f := func(x map[string]string) {
			matches := IsNil(x)
			assertTrue(t, matches)
		}
		f(nil)
	})

	t.Run("not_nil_map", func(t *testing.T) {
		f := func(x map[string]string) {
			matches := IsNil(x)
			assertFalse(t, matches)
		}
		f(make(map[string]string, 0))
	})

	t.Run("nil_ptr", func(t *testing.T) {
		f := func(x *int) {
			matches := IsNil(x)
			assertTrue(t, matches)
		}
		f(nil)
	})

	t.Run("not_nil_ptr", func(t *testing.T) {
		ptr := new(int)
		f := func(x *int) {
			matches := IsNil(x)
			assertFalse(t, matches)
		}
		f(ptr)
	})
}

func assertTrue(t *testing.T, value bool) {
	t.Helper()
	if !value {
		t.Errorf("want true; got false")
	}
}

func assertFalse(t *testing.T, value bool) {
	t.Helper()
	if value {
		t.Errorf("want false; got true")
	}
}
