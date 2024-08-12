package searchingalgorithms_test

import (
	"testing"

	searchingalgorithms "github.com/agustfricke/dsa-go/searching_algorithms"
)

func TestLinearSearch(t *testing.T) {
	t.Run("search empty array", func(t *testing.T) {
		arr := []int{}
		target := 69
		index := searchingalgorithms.LinearSearch(arr, target)
		if index != -1 {
			t.Fatalf("expected value -1, but got %v", index)
		}
	})

	t.Run("search not found", func(t *testing.T) {
		arr := []int{420, 73, 333}
		target := 69
		index := searchingalgorithms.LinearSearch(arr, target)
		if index != -1 {
			t.Fatalf("expected value -1, but got %v", index)
		}
	})

	t.Run("search success found", func(t *testing.T) {
		arr := []int{420, 73, 333, 34, 19, 20, 69, 37, 19}
		target := 69
		index := searchingalgorithms.LinearSearch(arr, target)
		if index != 6 {
			t.Fatalf("expected value -1, but got %v", index)
		}
	})
}
