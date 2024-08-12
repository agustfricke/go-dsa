package searchingalgorithms_test

import (
	searchingalgorithms "github.com/agustfricke/dsa-go/searching_algorithms"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	eArr := []int{}
	testCases := []struct {
		name     string
		array    []int
		target   int
		expected int
	}{
		{"Element at the beginning", arr, 1, 0},
		{"Element at the end", arr, 19, 9},
		{"Element in the middle", arr, 11, 5},
		{"Element not present (lower)", arr, 0, -1},
		{"Element not present (higher)", arr, 20, -1},
		{"Element not present (between values)", arr, 6, -1},
		{"Search first repeated element", arr, 3, 1},
		{"Search last repeated element", arr, 17, 8},
		{"Search in empty arr", eArr, 17, 8},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := searchingalgorithms.BinarySearch(arr, tc.target)
			if result != tc.expected {
				t.Errorf("For case '%s': expected %d, but got %d", tc.name, tc.expected, result)
			}
		})
	}
}
