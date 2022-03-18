package internal

import (
	"fmt"
	"testing"
)

func TestIsEqual(t *testing.T) {

	var testCases = []struct {
		actual          interface{}
		expected        interface{}
		expectedToMatch bool
	}{
		{actual: nil, expected: nil, expectedToMatch: true},
		{actual: 1, expected: nil, expectedToMatch: false},
		{actual: nil, expected: 2, expectedToMatch: false},
		{actual: 1, expected: 1, expectedToMatch: true},

		{actual: true, expected: true, expectedToMatch: true},
		{actual: false, expected: true, expectedToMatch: false},
		{actual: true, expected: false, expectedToMatch: false},
		{actual: false, expected: false, expectedToMatch: true},

		{actual: 10, expected: 42.24, expectedToMatch: false},

		{actual: float32(13.37), expected: float32(42.24), expectedToMatch: false},
		{actual: float32(13.37), expected: float32(13.37), expectedToMatch: true},

		{actual: float64(13.37), expected: float64(42.24), expectedToMatch: false},
		{actual: float64(13.37), expected: float64(13.37), expectedToMatch: true},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Case %d -- %v", idx+1, testCase), func(t *testing.T) {
			if IsEqual(testCase.actual, testCase.expected) != testCase.expectedToMatch {
				t.Errorf("Case %d failed -- %v", idx+1, testCase)
			}
		})
	}
}
