/*
Search Operation:
In an unsorted array, the search operation can be performed by linear
traversal from the first element to the last element.
*/
package array

// Function to implement search operation
func FindElement(arr []int, key int) int {
	for i, value := range arr {
		if value == key {
			return i
		}
	}
	// If the key is not found
	return -1
}
