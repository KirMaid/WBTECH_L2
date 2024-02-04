package unpack

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{"Simple", "a4bc2d5e", "aaaabccddddde", false},
		{"No repeats", "abcd", "abcd", false},
		{"Invalid", "45", "", true},
		{"Escape sequences", "qwe\\4\\5", "qwe45", false},
		{"Multiple digits", "qwe\\45", "qwe44444", false},
		{"Double escape", "qwe\\\\5", "qwe\\\\\\", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Unpack(tc.input)
			if (err != nil) != tc.hasError {
				t.Errorf("Expected error: %v, got error: %v", tc.hasError, err)
			}
			if result != tc.expected {
				t.Errorf("Expected: %s, got: %s", tc.expected, result)
			}
		})
	}
}
