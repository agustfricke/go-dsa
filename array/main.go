package main

import (
    "fmt"
    "strconv"
)

func main() {
    var numeros [5]int
    for i := range numeros {
        numeros[i] = i + 1
    }

    fmt.Println("Array:", numeros)
    for i := 0; i < len(numeros); i++ {
        fmt.Printf("Dirección del elemento %d: %p\n", i, &numeros[i])
    }
    
    // Obtener y convertir la dirección base
    baseAddr, _ := strconv.ParseUint(fmt.Sprintf("%p", &numeros[0]), 0, 64)
    fmt.Printf("Dirección base (decimal): %d\n", baseAddr)

    // Mostrar las direcciones de los elementos en decimal
    for i := 0; i < len(numeros); i++ {
        // addr, _ := strconv.ParseInt(fmt.Sprintf("%p", &numeros[i]), 0, 64)
        fmt.Printf("elemento %d de hexadecimal (%d)\n", i, &numeros[i])
    }
}
