package bubblesort_test

import (
	"testing"

	bubblesort "github.com/agustfricke/dsa-go/bubble_sort"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Bubble sort success", func(t *testing.T) {
		arr := []int{1, 3, 7, 4, 2}
		expected := []int{1, 2, 3, 4, 7}
		sortedArr := bubblesort.BubbleSort(arr)
		for i, v := range sortedArr {
			if v != expected[i] {
				t.Fatalf("Expected %v at index %v but got %v", expected[i], i, v)
			}
		}
	})
	t.Run("Bubble sort in sorted array", func(t *testing.T) {
		arr := []int{1, 2, 4, 8, 9}
		expected := []int{1, 2, 4, 8, 9}
		sortedArr := bubblesort.BubbleSort(arr)
		for i, v := range sortedArr {
			if v != expected[i] {
				t.Fatalf("Expected %v at index %v but got %v", expected[i], i, v)
			}
		}
	})
	t.Run("Bubble sort with duplicated element in the array", func(t *testing.T) {
		arr := []int{1, 2, 4, 8, 4}
		expected := []int{1, 2, 4, 4, 8}
		sortedArr := bubblesort.BubbleSort(arr)
		for i, v := range sortedArr {
			if v != expected[i] {
				t.Fatalf("Expected %v at index %v but got %v", expected[i], i, v)
			}
		}
	})
	t.Run("Bubble sort in empty array", func(t *testing.T) {
		arr := []int{}
		expected := []int{}
		sortedArr := bubblesort.BubbleSort(arr)
		for i, v := range sortedArr {
			if v != expected[i] {
				t.Fatalf("Expected %v at index %v but got %v", expected[i], i, v)
			}
		}
	})
}
