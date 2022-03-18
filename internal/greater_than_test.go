package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestIsGreaterThan(t *testing.T) {

	var testCases = []struct {
		actual          interface{}
		expected        interface{}
		expectedToMatch bool
	}{
		{actual: 0, expected: 0, expectedToMatch: false},
		{actual: 1, expected: 0, expectedToMatch: true},
		{actual: 0, expected: 1, expectedToMatch: false},

		{actual: int(1), expected: int(2), expectedToMatch: false},
		{actual: int(2), expected: int(1), expectedToMatch: true},
		{actual: int(1), expected: int(1), expectedToMatch: false},

		{actual: int8(1), expected: int8(2), expectedToMatch: false},
		{actual: int8(2), expected: int8(1), expectedToMatch: true},
		{actual: int8(1), expected: int8(1), expectedToMatch: false},

		{actual: int16(1), expected: int16(2), expectedToMatch: false},
		{actual: int16(2), expected: int16(1), expectedToMatch: true},
		{actual: int16(1), expected: int16(1), expectedToMatch: false},

		{actual: int32(1), expected: int32(2), expectedToMatch: false},
		{actual: int32(2), expected: int32(1), expectedToMatch: true},
		{actual: int32(1), expected: int32(1), expectedToMatch: false},

		{actual: int64(1), expected: int64(2), expectedToMatch: false},
		{actual: int64(2), expected: int64(1), expectedToMatch: true},
		{actual: int64(1), expected: int64(1), expectedToMatch: false},

		{actual: uint(1), expected: uint(2), expectedToMatch: false},
		{actual: uint(2), expected: uint(1), expectedToMatch: true},
		{actual: uint(1), expected: uint(1), expectedToMatch: false},

		{actual: uint8(1), expected: uint8(2), expectedToMatch: false},
		{actual: uint8(2), expected: uint8(1), expectedToMatch: true},
		{actual: uint8(1), expected: uint8(1), expectedToMatch: false},

		{actual: uint16(1), expected: uint16(2), expectedToMatch: false},
		{actual: uint16(2), expected: uint16(1), expectedToMatch: true},
		{actual: uint16(1), expected: uint16(1), expectedToMatch: false},

		{actual: uint32(1), expected: uint32(2), expectedToMatch: false},
		{actual: uint32(2), expected: uint32(1), expectedToMatch: true},
		{actual: uint32(1), expected: uint32(1), expectedToMatch: false},

		{actual: uint64(1), expected: uint64(2), expectedToMatch: false},
		{actual: uint64(2), expected: uint64(1), expectedToMatch: true},
		{actual: uint64(1), expected: uint64(1), expectedToMatch: false},

		{actual: float32(1.12), expected: float32(1.11), expectedToMatch: true},
		{actual: float32(1.11), expected: float32(1.12), expectedToMatch: false},
		{actual: float32(1.11), expected: float32(1.11), expectedToMatch: false},

		{actual: float64(1.12), expected: float64(1.11), expectedToMatch: true},
		{actual: float64(1.11), expected: float64(1.12), expectedToMatch: false},
		{actual: float64(1.11), expected: float64(1.11), expectedToMatch: false},

		{actual: "xyz", expected: "abc", expectedToMatch: true},
		{actual: "abc", expected: "xyz", expectedToMatch: false},
		{actual: "abc", expected: "abc", expectedToMatch: false},

		{actual: false, expected: false, expectedToMatch: false},
		{actual: true, expected: false, expectedToMatch: true},
		{actual: false, expected: true, expectedToMatch: false},
		{actual: true, expected: true, expectedToMatch: false},

		{actual: 1 * time.Second, expected: 1 * time.Minute, expectedToMatch: false},
		{actual: 1 * time.Minute, expected: 1 * time.Minute, expectedToMatch: false},
		{actual: 1 * time.Minute, expected: 1 * time.Second, expectedToMatch: true},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Case %d -- %v", idx+1, testCase), func(t *testing.T) {
			if IsGreaterThan(testCase.actual, testCase.expected) != testCase.expectedToMatch {
				t.Errorf("Case %d failed -- %v", idx+1, testCase)
			}
		})
	}
}

func TestIsGreaterThanEqual(t *testing.T) {

	var testCases = []struct {
		actual          interface{}
		expected        interface{}
		expectedToMatch bool
	}{
		{actual: 0, expected: 0, expectedToMatch: true},
		{actual: 1, expected: 0, expectedToMatch: true},
		{actual: 0, expected: 1, expectedToMatch: false},

		{actual: int(1), expected: int(2), expectedToMatch: false},
		{actual: int(2), expected: int(1), expectedToMatch: true},
		{actual: int(2), expected: int(2), expectedToMatch: true},

		{actual: int8(1), expected: int8(2), expectedToMatch: false},
		{actual: int8(2), expected: int8(1), expectedToMatch: true},
		{actual: int8(2), expected: int8(2), expectedToMatch: true},

		{actual: int16(1), expected: int16(2), expectedToMatch: false},
		{actual: int16(2), expected: int16(1), expectedToMatch: true},
		{actual: int16(2), expected: int16(2), expectedToMatch: true},

		{actual: int32(1), expected: int32(2), expectedToMatch: false},
		{actual: int32(2), expected: int32(1), expectedToMatch: true},
		{actual: int32(2), expected: int32(2), expectedToMatch: true},

		{actual: int64(1), expected: int64(2), expectedToMatch: false},
		{actual: int64(2), expected: int64(1), expectedToMatch: true},
		{actual: int64(2), expected: int64(2), expectedToMatch: true},

		{actual: uint(1), expected: uint(2), expectedToMatch: false},
		{actual: uint(2), expected: uint(1), expectedToMatch: true},
		{actual: uint(2), expected: uint(2), expectedToMatch: true},

		{actual: uint8(1), expected: uint8(2), expectedToMatch: false},
		{actual: uint8(2), expected: uint8(1), expectedToMatch: true},
		{actual: uint8(2), expected: uint8(2), expectedToMatch: true},

		{actual: uint16(1), expected: uint16(2), expectedToMatch: false},
		{actual: uint16(2), expected: uint16(1), expectedToMatch: true},
		{actual: uint16(2), expected: uint16(2), expectedToMatch: true},

		{actual: uint32(1), expected: uint32(2), expectedToMatch: false},
		{actual: uint32(2), expected: uint32(1), expectedToMatch: true},
		{actual: uint32(2), expected: uint32(2), expectedToMatch: true},

		{actual: uint64(1), expected: uint64(2), expectedToMatch: false},
		{actual: uint64(2), expected: uint64(1), expectedToMatch: true},

		{actual: float32(1.12), expected: float32(1.11), expectedToMatch: true},
		{actual: float32(1.11), expected: float32(1.12), expectedToMatch: false},
		{actual: float32(1.11), expected: float32(1.11), expectedToMatch: true},

		{actual: float64(1.11), expected: float64(1.12), expectedToMatch: false},
		{actual: float64(1.12), expected: float64(1.11), expectedToMatch: true},
		{actual: float64(1.11), expected: float64(1.11), expectedToMatch: true},

		{actual: "xyz", expected: "abc", expectedToMatch: true},
		{actual: "abc", expected: "xyz", expectedToMatch: false},
		{actual: "abc", expected: "abc", expectedToMatch: true},

		{actual: false, expected: false, expectedToMatch: true},
		{actual: false, expected: true, expectedToMatch: false},
		{actual: true, expected: false, expectedToMatch: true},
		{actual: true, expected: true, expectedToMatch: true},

		{actual: 1 * time.Second, expected: 1 * time.Minute, expectedToMatch: false},
		{actual: 1 * time.Minute, expected: 1 * time.Minute, expectedToMatch: true},
		{actual: 1 * time.Minute, expected: 1 * time.Second, expectedToMatch: true},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Case %d -- %v", idx+1, testCase), func(t *testing.T) {
			if IsGreaterThanEqual(testCase.actual, testCase.expected) != testCase.expectedToMatch {
				t.Errorf("Case %d failed -- %v", idx+1, testCase)
			}
		})
	}
}
