package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name         string
		input        []string
		expected_res int
	}{
		{
			name:         "1",
			input:        []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"},
			expected_res: 161,
		},
		{
			name:         "2",
			input:        []string{"mul(911,800)", "mul(734,19)", "mul(520,383)", "mul(700,211)"},
			expected_res: 1089606,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual_res := add(tt.input, false)
			if actual_res != tt.expected_res {
				t.Errorf("add([]string) = %v, want %v", actual_res, tt.expected_res)
			}
		})
	}
}
