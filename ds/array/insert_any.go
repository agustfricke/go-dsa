/*
Insert operation in an array at any position can be performed by
shifting elements to the right, which are on the right side of
the required position
*/
package array

import "fmt"

// InsertElement inserts an element at a specific position in the slice.
func InsertElement(arr []int, x int, pos int) []int {
	// Check if the position is within the bounds of the slice
	if pos < 0 || pos > len(arr) {
		fmt.Println("Position out of bounds")
		return arr
	}

	// Create a new slice with space for one more element
	newArr := make([]int, len(arr)+1)

	// Copy elements before the insertion point
	copy(newArr[:pos], arr[:pos])

	// Insert the new element
	newArr[pos] = x

	// Copy elements after the insertion point
	copy(newArr[pos+1:], arr[pos:])

	return newArr
}
