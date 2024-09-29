package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected int
	}{
		{"Valid English", []byte("Hello"), 5},
		{"Single invalid byte", []byte{0xc2}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetUTFLength(tt.input)
			if err != nil {
				if err.Error() != "invalid utf8" {
					t.Errorf("GetUTFLength(%v) error = %v, want \"invalid utf8\"", tt.input, err)
				}
			} else if result != tt.expected {
				t.Errorf("GetUTFLength(%v) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
