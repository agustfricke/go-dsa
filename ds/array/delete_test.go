package array_test

import (
	"testing"

	"github.com/agustfricke/dsa-go/ds/array"
)

func TestDeleteElement(t *testing.T) {
	tests := []struct {
		arr  []int
		key  int
		want []int
	}{
		{
			arr:  []int{10, 50, 30, 40, 20},
			key:  30,
			want: []int{10, 50, 40, 20},
		},
		{
			arr:  []int{10, 50, 30, 40, 20},
			key:  10,
			want: []int{50, 30, 40, 20},
		},
		{
			arr:  []int{10, 50, 30, 40, 20},
			key:  20,
			want: []int{10, 50, 30, 40},
		},
		{
			arr:  []int{10, 50, 30, 40, 20},
			key:  100, // Key not present
			want: []int{10, 50, 30, 40, 20},
		},
		{
			arr:  []int{},
			key:  10, // Key not present in an empty slice
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := array.DeleteElement(tt.arr, tt.key)

			// Check if the result slice matches the expected result
			if !array.Equal(got, tt.want) {
				t.Errorf("deleteElement(%v, %d) = %v; want %v",
					tt.arr, tt.key, got, tt.want)
			}
		})
	}
}
