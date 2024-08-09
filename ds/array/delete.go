/*
In the delete operation, the element to be deleted is searched using
the linear search, and then the delete operation is performed followed
by shifting the elements.
*/
package array

import "fmt"

// FindElement searches for a key in the slice and returns its index.
// Returns -1 if the key is not found.
func findElement(arr []int, key int) int {
	for i, value := range arr {
		if value == key {
			return i
		}
	}
	return -1
}

// DeleteElement deletes an element with the given key from the slice.
// It returns the new length of the slice after deletion.
func DeleteElement(arr []int, key int) []int {
	// Find position of element to be deleted
	pos := findElement(arr, key)

	if pos == -1 {
		fmt.Println("Element not found")
		return arr
	}

	// Deleting element by shifting elements left
	arr = append(arr[:pos], arr[pos+1:]...)

	return arr
}
