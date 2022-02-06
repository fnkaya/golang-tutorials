package assignment

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	tcs := []struct {
		given            [2]uint32
		expectedSum      uint32
		expectedOverflow bool
	}{
		{[2]uint32{math.MaxUint32, 1}, 0, true},
		{[2]uint32{1, 1}, 2, false},
		{[2]uint32{42, 2701}, 2743, false},
		{[2]uint32{42, math.MaxUint32}, 41, true},
		{[2]uint32{4294967290, 5}, 4294967295, false},
		{[2]uint32{4294967290, 6}, 0, true},
		{[2]uint32{4294967290, 10}, 4, true},
	}

	for _, tc := range tcs {
		sum, overflow := AddUint32(tc.given[0], tc.given[1])
		assert.Equal(t, tc.expectedSum, sum)
		assert.Equal(t, tc.expectedOverflow, overflow)
	}
}

func TestCeilNumber(t *testing.T) {
	tcs := []struct {
		given    float64
		expected float64
	}{
		{42.42, 42.50},
		{42, 42},
		{42.01, 42.25},
		{42.24, 42.25},
		{42.25, 42.25},
		{42.26, 42.50},
		{42.55, 42.75},
		{42.75, 42.75},
		{42.76, 43},
		{42.99, 43},
		{43.13, 43.25},
	}

	for _, tc := range tcs {
		result := CeilNumber(tc.given)
		assert.Equal(t, tc.expected, result)
	}
}

func TestAlphabetSoup(t *testing.T) {
	tcs := []struct {
		given    string
		expected string
	}{
		{"hello", "ehllo"},
		{"", ""},
		{"h", "h"},
		{"ab", "ab"},
		{"ba", "ab"},
		{"bac", "abc"},
		{"cba", "abc"},
	}

	for _, tc := range tcs {
		output := AlphabetSoup(tc.given)
		assert.Equal(t, tc.expected, output)
	}
}

func TestStringMask(t *testing.T) {
	tcs := []struct {
		givenValue   string
		givenUnmaskC uint
		expected     string
	}{
		{"!mysecret*", 2, "!m********"},
		{"", 100, "*"},
		{"a", 1, "*"},
		{"string", 0, "******"},
		{"string", 3, "str***"},
		{"string", 5, "strin*"},
		{"string", 6, "******"},
		{"string", 7, "******"},
		{"s*r*n*", 3, "s*r***"},
	}

	for _, tc := range tcs {
		output := StringMask(tc.givenValue, tc.givenUnmaskC)
		assert.Equal(t, tc.expected, output)
	}
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	tcs := []struct {
		given    [2]string
		expected interface{}
		// I prefer to change expected value as interface
	}{
		{[2]string{"hellocat", words}, [2]string{"hello", "cat"}},
		{[2]string{"catbat", words}, [2]string{"cat", "bat"}},
		{[2]string{"yellowapple", words}, [2]string{"yellow", "apple"}},
		{[2]string{"", words}, "not possible"},
		{[2]string{"notcat", words}, "not possible"},
		{[2]string{"bootcamprocks!", words}, "not possible"},
	}

	for _, tc := range tcs {
		output := WordSplit(tc.given)
		assert.ObjectsAreEqualValues(tc.expected, output)
	}
}

func TestVariadicSet(t *testing.T) {
	tcs := []struct {
		given    []interface{}
		expected []interface{}
	}{
		{
			[]interface{}{4, 2, 5, 4, 2, 4},
			[]interface{}{4, 2, 5}},
		{
			[]interface{}{"bootcamp", "rocks!", "really", "rocks!"},
			[]interface{}{"bootcamp", "rocks!", "really"}},
		{
			[]interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"},
			[]interface{}{1, uint32(1), "first", 2, uint32(2), "second"}},
	}

	for _, tc := range tcs {
		output := VariadicSet(tc.given...)
		assert.ElementsMatch(t, tc.expected, output)
	}
}
