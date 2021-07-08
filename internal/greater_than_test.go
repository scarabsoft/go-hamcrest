package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestIsGreaterThan(t *testing.T) {

	var testCases = []struct {
		actual          interface{}
		given           interface{}
		expectedToMatch bool
	}{
		{actual: 0, given: 0, expectedToMatch: false},
		{actual: 1, given: 0, expectedToMatch: true},
		{actual: 0, given: 1, expectedToMatch: false},

		{actual: int(1), given: int(2), expectedToMatch: false},
		{actual: int(2), given: int(1), expectedToMatch: true},
		{actual: int(1), given: int(1), expectedToMatch: false},

		{actual: int8(1), given: int8(2), expectedToMatch: false},
		{actual: int8(2), given: int8(1), expectedToMatch: true},
		{actual: int8(1), given: int8(1), expectedToMatch: false},

		{actual: int16(1), given: int16(2), expectedToMatch: false},
		{actual: int16(2), given: int16(1), expectedToMatch: true},
		{actual: int16(1), given: int16(1), expectedToMatch: false},

		{actual: int32(1), given: int32(2), expectedToMatch: false},
		{actual: int32(2), given: int32(1), expectedToMatch: true},
		{actual: int32(1), given: int32(1), expectedToMatch: false},

		{actual: int64(1), given: int64(2), expectedToMatch: false},
		{actual: int64(2), given: int64(1), expectedToMatch: true},
		{actual: int64(1), given: int64(1), expectedToMatch: false},

		{actual: uint(1), given: uint(2), expectedToMatch: false},
		{actual: uint(2), given: uint(1), expectedToMatch: true},
		{actual: uint(1), given: uint(1), expectedToMatch: false},

		{actual: uint8(1), given: uint8(2), expectedToMatch: false},
		{actual: uint8(2), given: uint8(1), expectedToMatch: true},
		{actual: uint8(1), given: uint8(1), expectedToMatch: false},

		{actual: uint16(1), given: uint16(2), expectedToMatch: false},
		{actual: uint16(2), given: uint16(1), expectedToMatch: true},
		{actual: uint16(1), given: uint16(1), expectedToMatch: false},

		{actual: uint32(1), given: uint32(2), expectedToMatch: false},
		{actual: uint32(2), given: uint32(1), expectedToMatch: true},
		{actual: uint32(1), given: uint32(1), expectedToMatch: false},

		{actual: uint64(1), given: uint64(2), expectedToMatch: false},
		{actual: uint64(2), given: uint64(1), expectedToMatch: true},
		{actual: uint64(1), given: uint64(1), expectedToMatch: false},

		{actual: float32(1.12), given: float32(1.11), expectedToMatch: true},
		{actual: float32(1.11), given: float32(1.12), expectedToMatch: false},
		{actual: float32(1.11), given: float32(1.11), expectedToMatch: false},

		{actual: float64(1.12), given: float64(1.11), expectedToMatch: true},
		{actual: float64(1.11), given: float64(1.12), expectedToMatch: false},
		{actual: float64(1.11), given: float64(1.11), expectedToMatch: false},

		{actual: "xyz", given: "abc", expectedToMatch: true},
		{actual: "abc", given: "xyz", expectedToMatch: false},
		{actual: "abc", given: "abc", expectedToMatch: false},

		{actual: false, given: false, expectedToMatch: false},
		{actual: true, given: false, expectedToMatch: true},
		{actual: false, given: true, expectedToMatch: false},
		{actual: true, given: true, expectedToMatch: false},

		{actual: 1 * time.Second, given: 1 * time.Minute, expectedToMatch: false},
		{actual: 1 * time.Minute, given: 1 * time.Minute, expectedToMatch: false},
		{actual: 1 * time.Minute, given: 1 * time.Second, expectedToMatch: true},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Case %d -- %v", idx+1, testCase), func(t *testing.T) {
			if IsGreaterThan(testCase.actual, testCase.given) != testCase.expectedToMatch {
				t.Errorf("Case %d failed -- %v", idx+1, testCase)
			}
		})
	}
}

func TestIsGreaterThanEqual(t *testing.T) {

	var testCases = []struct {
		actual          interface{}
		given           interface{}
		expectedToMatch bool
	}{
		{actual: 0, given: 0, expectedToMatch: true},
		{actual: 1, given: 0, expectedToMatch: true},
		{actual: 0, given: 1, expectedToMatch: false},

		{actual: int(1), given: int(2), expectedToMatch: false},
		{actual: int(2), given: int(1), expectedToMatch: true},
		{actual: int(2), given: int(2), expectedToMatch: true},

		{actual: int8(1), given: int8(2), expectedToMatch: false},
		{actual: int8(2), given: int8(1), expectedToMatch: true},
		{actual: int8(2), given: int8(2), expectedToMatch: true},

		{actual: int16(1), given: int16(2), expectedToMatch: false},
		{actual: int16(2), given: int16(1), expectedToMatch: true},
		{actual: int16(2), given: int16(2), expectedToMatch: true},

		{actual: int32(1), given: int32(2), expectedToMatch: false},
		{actual: int32(2), given: int32(1), expectedToMatch: true},
		{actual: int32(2), given: int32(2), expectedToMatch: true},

		{actual: int64(1), given: int64(2), expectedToMatch: false},
		{actual: int64(2), given: int64(1), expectedToMatch: true},
		{actual: int64(2), given: int64(2), expectedToMatch: true},

		{actual: uint(1), given: uint(2), expectedToMatch: false},
		{actual: uint(2), given: uint(1), expectedToMatch: true},
		{actual: uint(2), given: uint(2), expectedToMatch: true},

		{actual: uint8(1), given: uint8(2), expectedToMatch: false},
		{actual: uint8(2), given: uint8(1), expectedToMatch: true},
		{actual: uint8(2), given: uint8(2), expectedToMatch: true},

		{actual: uint16(1), given: uint16(2), expectedToMatch: false},
		{actual: uint16(2), given: uint16(1), expectedToMatch: true},
		{actual: uint16(2), given: uint16(2), expectedToMatch: true},

		{actual: uint32(1), given: uint32(2), expectedToMatch: false},
		{actual: uint32(2), given: uint32(1), expectedToMatch: true},
		{actual: uint32(2), given: uint32(2), expectedToMatch: true},

		{actual: uint64(1), given: uint64(2), expectedToMatch: false},
		{actual: uint64(2), given: uint64(1), expectedToMatch: true},

		{actual: float32(1.12), given: float32(1.11), expectedToMatch: true},
		{actual: float32(1.11), given: float32(1.12), expectedToMatch: false},
		{actual: float32(1.11), given: float32(1.11), expectedToMatch: true},

		{actual: float64(1.11), given: float64(1.12), expectedToMatch: false},
		{actual: float64(1.12), given: float64(1.11), expectedToMatch: true},
		{actual: float64(1.11), given: float64(1.11), expectedToMatch: true},

		{actual: "xyz", given: "abc", expectedToMatch: true},
		{actual: "abc", given: "xyz", expectedToMatch: false},
		{actual: "abc", given: "abc", expectedToMatch: true},

		{actual: false, given: false, expectedToMatch: true},
		{actual: false, given: true, expectedToMatch: false},
		{actual: true, given: false, expectedToMatch: true},
		{actual: true, given: true, expectedToMatch: true},

		{actual: 1 * time.Second, given: 1 * time.Minute, expectedToMatch: false},
		{actual: 1 * time.Minute, given: 1 * time.Minute, expectedToMatch: true},
		{actual: 1 * time.Minute, given: 1 * time.Second, expectedToMatch: true},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Case %d -- %v", idx+1, testCase), func(t *testing.T) {
			if IsGreaterThanEqual(testCase.actual, testCase.given) != testCase.expectedToMatch {
				t.Errorf("Case %d failed -- %v", idx+1, testCase)
			}
		})
	}
}
