/*
In an unsorted array, the insert operation is faster 
as compared to a sorted array because we don’t have to 
care about the position at which the element is to be placed.
*/
package array

// Inserts a key into the slice if there is capacity
// It returns the new length of the slice after insertion
func InsertSorted(arr []int, n int, key int, capacity int) int {
    // Cannot insert more elements if n is already more than or 
    // equal to capacity
    if n >= capacity {
        return n
    }

    // Insert the key
    arr[n] = key

    return n + 1
}
