package internal

import (
	"fmt"
	"testing"
)

func testChannel(size int) chan int {
	result := make(chan int, size)
	for idx := 0; idx < size; idx++ {
		result <- idx
	}
	return result
}

func TestLength(t *testing.T) {
	var someString = "someString"
	var someInt = 42

	var testCases = []struct {
		actual interface{}
		length int64
		error  string
	}{
		{actual: [0]string{}, length: 0, error: ""},
		{actual: [1]string{"given"}, length: 1, error: ""},
		{actual: &[0]string{}, length: 0, error: ""},

		{actual: []string{}, length: 0, error: ""},
		{actual: []string{"given"}, length: 1, error: ""},
		{actual: &[]string{}, length: 0, error: ""},

		{actual: map[string]string{}, length: 0, error: ""},
		{actual: map[string]string{"givenKey": "givenValue"}, length: 1, error: ""},
		{actual: &map[string]string{}, length: 0, error: ""},

		{actual: "", length: 0, error: ""},
		{actual: someString, length: 10, error: ""},
		{actual: &someString, length: 10, error: ""},

		{actual: testChannel(0), length: 0, error: ""},
		{actual: testChannel(3), length: 3, error: ""},

		{actual: &someInt, length: -1, error: "actual pointer not one of [array,map,slice,string,ptr]"},
		{actual: someInt, length: -1, error: "actual not one of [array,chan,map,slice,string,ptr]"},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Case %d -- %v", idx+1, testCase), func(t *testing.T) {
			length, err := Length(testCase.actual)

			if err != nil && !IsEqual(testCase.error, err.Error()) {
				t.Errorf("Case %d failed -- %v with error miss match; want %s; got %s", idx+1, testCase, testCase.error, err)
			}

			if !IsEqual(testCase.length, length) {
				t.Errorf("Case %d failed -- %v with length miss match; want %d; got %d", idx+1, testCase, testCase.length, length)
			}
		})
	}
}

func TestHasItem(t *testing.T) {

	type someObject struct {
		value int
	}

	var testCases = []struct {
		actual  interface{}
		item    interface{}
		hasItem bool
		error   string
	}{
		{actual: [0]int{}, item: 0, hasItem: false, error: ""},
		{actual: [1]int{1}, item: 1, hasItem: true, error: ""},
		{actual: [1]int{1}, item: 1.2, hasItem: false, error: ""},
		{actual: [1]int{1}, item: "1", hasItem: false, error: ""},
		{actual: []string{"abc", "def"}, item: "xyz", hasItem: false, error: ""},
		{actual: []string{"abc", "def", "xyz"}, item: "xyz", hasItem: true, error: ""},
		{actual: &[]string{"abc", "def", "xyz"}, item: "xyz", hasItem: true, error: ""},
		{actual: []struct{}{}, item: nil, hasItem: false, error: ""},
		{actual: []*struct{}{nil}, item: nil, hasItem: true, error: ""},
		{actual: []someObject{{1}, {2}}, item: someObject{3}, hasItem: false, error: ""},
		{actual: &[]someObject{{1}, {2}}, item: someObject{2}, hasItem: true, error: ""},
		{actual: []*someObject{{1}, {2}}, item: &someObject{2}, hasItem: true, error: ""},
		{actual: testChannel(0), item: 1, hasItem: false, error: ""},
		{actual: testChannel(5), item: 4, hasItem: true, error: ""},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Case %d -- %v", idx+1, testCase), func(t *testing.T) {
			hasItem, err := HasItem(testCase.actual, testCase.item)

			if err != nil && !IsEqual(testCase.error, err.Error()) {
				t.Errorf("Case %d failed -- %v with error miss match; want %s; got %s", idx+1, testCase, testCase.error, err)
			}

			if !IsEqual(testCase.hasItem, hasItem) {
				t.Errorf("Case %d failed -- %v with hasItem miss match; want %t; got %t", idx+1, testCase, testCase.hasItem, hasItem)
			}
		})
	}

	// inspecting a channel does not seems to be trivial, as we have to consume elements in order to compare them
	// but as we want to use the HasItem method for HasItems as well, we have to make sure that we dont alter the channel
	t.Run("special_case_channel", func(t *testing.T) {
		givenChannel := testChannel(10)

		assertTrue(t, len(givenChannel) == 10)
		result, err := HasItem(givenChannel, 5)
		assertTrue(t, err == nil)
		assertTrue(t, result)
		assertTrue(t, len(givenChannel) == 10)

		for i := 0; i < 10; i++ {
			v := <-givenChannel
			assertTrue(t, i == v)
		}
	})
}
