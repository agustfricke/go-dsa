package array_test

import (
	"testing"

	"github.com/agustfricke/dsa-go/ds/array"
)

func TestInsertSorted(t *testing.T) {
	tests := []struct {
		arr      []int
		n        int
		key      int
		capacity int
		want     []int
		wantLen  int
	}{
		{
			arr:      []int{12, 16, 20, 40, 50, 70},
			n:        6,
			key:      26,
			capacity: 20,
			want:     []int{12, 16, 20, 40, 50, 70, 26},
			wantLen:  7,
		},
		{
			arr:      []int{12, 16, 20, 40, 50, 70},
			n:        6,
			key:      5,
			capacity: 6,
			want:     []int{12, 16, 20, 40, 50, 70},
			wantLen:  6,
		},
		{
			arr:      []int{12, 16, 20},
			n:        3,
			key:      25,
			capacity: 5,
			want:     []int{12, 16, 20, 25},
			wantLen:  4,
		},
		{
			arr:      []int{12, 16, 20},
			n:        3,
			key:      18,
			capacity: 3,
			want:     []int{12, 16, 20},
			wantLen:  3,
		},
		{
			arr:      []int{},
			n:        0,
			key:      10,
			capacity: 10,
			want:     []int{10},
			wantLen:  1,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			// Make a copy of the slice to test the function
			arr := make([]int, len(tt.arr)+1) // Create a slice with one extra space for insertion
			copy(arr, tt.arr)

			gotLen := array.InsertSorted(arr, tt.n, tt.key, tt.capacity)

			if gotLen != tt.wantLen {
				t.Errorf("InsertSorted() = %d; want %d", gotLen, tt.wantLen)
			}

			if !array.Equal(arr[:gotLen], tt.want) {
				t.Errorf("InsertSorted() = %v; want %v", arr[:gotLen], tt.want)
			}
		})
	}
}
