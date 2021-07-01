package internal

import (
	"fmt"
	"testing"
)

func TestIsEqual(t *testing.T) {

	var testCases = []struct {
		actual          interface{}
		given           interface{}
		expectedToMatch bool
	}{
		{actual: nil, given: nil, expectedToMatch: true},
		{actual: 1, given: nil, expectedToMatch: false},
		{actual: nil, given: 2, expectedToMatch: false},
		{actual: 1, given: 1, expectedToMatch: true},

		{actual: true, given: true, expectedToMatch: true},
		{actual: false, given: true, expectedToMatch: false},
		{actual: true, given: false, expectedToMatch: false},
		{actual: false, given: false, expectedToMatch: true},

		{actual: 10, given: 42.24, expectedToMatch: false},

		{actual: float32(13.37), given: float32(42.24), expectedToMatch: false},
		{actual: float32(13.37), given: float32(13.37), expectedToMatch: true},

		{actual: float64(13.37), given: float64(42.24), expectedToMatch: false},
		{actual: float64(13.37), given: float64(13.37), expectedToMatch: true},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Case %d -- %v", idx+1, testCase), func(t *testing.T) {
			if IsEqual(testCase.actual, testCase.given) != testCase.expectedToMatch {
				t.Errorf("Case %d failed -- %v", idx+1, testCase)
			}
		})
	}
}
