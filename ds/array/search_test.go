package array_test

import (
	"testing"

	"github.com/agustfricke/dsa-go/ds/array"
)

func TestFindElement(t *testing.T) {
	tests := []struct {
		arr  []int
		key  int
		want int
	}{
		{[]int{12, 34, 10, 6, 40}, 40, 4},
		{[]int{12, 34, 10, 6, 40}, 10, 2},
		{[]int{12, 34, 10, 6, 40}, 5, -1},
		{[]int{12, 34, 10, 6, 40}, 12, 0},
		{[]int{}, 1, -1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := array.FindElement(tt.arr, tt.key)
			if got != tt.want {
				t.Errorf("FindElement(%v, %d) = %d; want %d",
					tt.arr, tt.key, got, tt.want)
			}
		})
	}
}
