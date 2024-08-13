package bubblesort

import "fmt"

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Outer iteration %d: %v\n", i+1, arr) // Print the array at the start of each outer iteration
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				fmt.Printf("Swapping %d and %d\n", arr[j], arr[j+1]) // Print the values before swapping
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				fmt.Printf("Current state: %v\n", arr) // Print the array after swapping
			}
		}
	}
	return arr
}
