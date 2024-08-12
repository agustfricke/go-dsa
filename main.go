package main

import (
	"fmt"

	searchingalgorithms "github.com/agustfricke/dsa-go/searching_algorithms"
)

func main() {
	arr := []int{11, 12, 22, 25, 34, 64, 90}
	target := 90

	result := searchingalgorithms.BinarySearch(arr, target)

	if result != -1 {
		fmt.Printf("Elemento %d encontrado en el índice %d\n", target, result)
	} else {
		fmt.Printf("Elemento %d no encontrado en el arreglo\n", target)
	}
}
