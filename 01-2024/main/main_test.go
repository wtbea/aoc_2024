package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name                string
		left                []int
		right               []int
		expected_distance   int
		expected_similarity int
	}{
		{
			name:                "1",
			left:                []int{3, 4, 2, 1, 3, 3}, // 1 2 3 3 3 4
			right:               []int{4, 3, 5, 3, 9, 3}, // 3 3 3 4 5 9
			expected_distance:   11,
			expected_similarity: 31,
		},
		{
			name:                "2",
			left:                []int{1, 4, 7, 6, 4}, // 1 4 4 6 7
			right:               []int{2, 8, 7, 7, 3}, // 2 3 7 7 8
			expected_distance:   7,
			expected_similarity: 14,
		},
		{
			name:                "3",
			left:                []int{5, 9, 5, 9, 8}, // 5 5 8 9 9
			right:               []int{8, 6, 5, 8, 7}, // 5 6 7 8 8
			expected_distance:   4,
			expected_similarity: 26,
		},
		{
			name:                "4",
			left:                []int{6, 3, 1, 4, 7}, // 1 3 4 6 7
			right:               []int{8, 3, 4, 5, 1}, // 1 3 4 5 8
			expected_distance:   2,
			expected_similarity: 8,
		},
		{
			name:                "Empty",
			left:                []int{},
			right:               []int{},
			expected_distance:   0,
			expected_similarity: 0,
		},
		{
			name:                "Single element",
			left:                []int{5},
			right:               []int{10},
			expected_distance:   5,
			expected_similarity: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual_distance := distance(tt.left, tt.right)
			if actual_distance != tt.expected_distance {
				t.Errorf("distance(%v, %v) = %d; want %d", tt.left, tt.right, actual_distance, tt.expected_distance)
			}

			actual_similarity := similarity(tt.left, tt.right)
			if actual_similarity != tt.expected_similarity {
				t.Errorf("similarity(%v, %v) = %d; want %d", tt.left, tt.right, actual_similarity, tt.expected_similarity)
			}
		})
	}
}
