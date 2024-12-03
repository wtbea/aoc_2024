package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name                     string
		list                     [][]int
		expected_safety          int
		expected_safety_dampener int
	}{
		{
			name:                     "1",
			list:                     [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}},
			expected_safety:          2,
			expected_safety_dampener: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual_safety := safety(tt.list, false)
			if actual_safety != tt.expected_safety {
				t.Errorf("safety(_, false) = %v, want %v", actual_safety, tt.expected_safety)
			}
			actual_safety_dampener := safety(tt.list, true)
			if actual_safety_dampener != tt.expected_safety_dampener {
				t.Errorf("safety(_, true) = %v, want %v", actual_safety_dampener, tt.expected_safety_dampener)
			}
		})
	}
}
