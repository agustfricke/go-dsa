package array_test

import (
	"testing"

	"github.com/agustfricke/dsa-go/ds/array"
)

func TestInsertElement(t *testing.T) {
	tests := []struct {
		arr  []int
		x    int
		pos  int
		want []int
	}{
		{
			arr:  []int{2, 4, 1, 8, 5},
			x:    10,
			pos:  2,
			want: []int{2, 4, 10, 1, 8, 5},
		},
		{
			arr:  []int{2, 4, 1, 8, 5},
			x:    3,
			pos:  0,
			want: []int{3, 2, 4, 1, 8, 5},
		},
		{
			arr:  []int{2, 4, 1, 8, 5},
			x:    7,
			pos:  5,
			want: []int{2, 4, 1, 8, 5, 7},
		},
		{
			arr:  []int{2, 4, 1, 8, 5},
			x:    9,
			pos:  10, // Position out of bounds
			want: []int{2, 4, 1, 8, 5},
		},
		{
			arr:  []int{},
			x:    10,
			pos:  0,
			want: []int{10},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := array.InsertElement(tt.arr, tt.x, tt.pos)

			// Check if the result slice matches the expected result
			if !array.Equal(got, tt.want) {
				t.Errorf("InsertElement(%v, %d, %d) = %v; want %v", tt.arr, tt.x, tt.pos, got, tt.want)
			}
		})
	}
}
