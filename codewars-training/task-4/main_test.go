package main

import (
	"testing"
)

// TestHighAndLow tests the HighAndLow function with various cases
func TestHighAndLow(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1 2 3 4 5", "5 1"},        // Simple ascending order
		{"1 9 3 4 -5", "9 -5"},      // Mixed positive and negative numbers
		{"42", "42 42"},             // Single number input
		{"-1 -2 -3 -4 -5", "-1 -5"}, // All negative numbers
		{"10 20 30 40 50", "50 10"}, // Simple case with larger numbers
	}

	for _, tt := range tests {
		result := HighAndLow(tt.input)
		if result != tt.expected {
			t.Errorf("HighAndLow(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
