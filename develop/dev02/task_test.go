package main

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errInvalidString},
		{"", "", nil},
	}

	for _, tc := range tests {
		result, err := unpack(tc.input)
		if result != tc.expected || !errors.Is(err, tc.err) {
			t.Errorf("unpack(%q) = %v, %v; want %v, %v", tc.input, result, err, tc.expected, tc.err)
		}
	}
}
