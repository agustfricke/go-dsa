package searchingalgorithms

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid // Elemento encontrado, retorna su índice
		}

		if arr[mid] < target {
			left = mid + 1 // Buscar en la mitad derecha
		} else {
			right = mid - 1 // Buscar en la mitad izquierda
		}
	}

	return -1 // Elemento no encontrado
}
