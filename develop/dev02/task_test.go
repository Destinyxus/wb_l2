package main

import "testing"

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
	}

	for _, tc := range testCases {
		result := unpackString(tc.input)
		if result != tc.output {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.output, result)
		}
	}
}
