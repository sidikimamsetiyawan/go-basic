package main

import (
	"testing"
)

func TestEncryptThis(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "104olle 119dlro"}, // Normal case
		{"", "String is empty."},           // Empty string case
		{"a", "97"},                        // Single character case
		{"ab", "97b"},                      // Two characters case
		{"abc", "97cb"},                    // Three characters case
		{"abcd", "97dcb"},                  // Four characters case
		{"abcde", "97edcb"},                // Five characters case
		{"a b c", "97 98 99"},              // Multiple single characters
		{"hello  ", "104olle"},             // Trailing spaces
		{"  hello", "104olle"},             // Leading spaces
		{"  ", ""},                         // Only spaces
		{"a b c d e", "97 98 99 100 101"},  // Multiple single characters
		{"test case", "116tse 99esa"},      // Normal case with two words
	}

	for _, test := range tests {
		result := EncryptThis(test.input)
		if result != test.expected {
			t.Errorf("EncryptThis(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
