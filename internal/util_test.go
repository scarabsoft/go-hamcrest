package internal

import (
	"testing"
)

func TestFormatTypeWithValue(t *testing.T) {

	t.Run("nil", func(t *testing.T) {
		r := FormatTypeWithValue(nil)
		assertTrue(t, r == "<nil>")
	})

	t.Run("simple_string", func(t *testing.T) {
		r := FormatTypeWithValue("simpleString")
		assertTrue(t, r == "string(simpleString)")
	})

}
