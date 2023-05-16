package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := 1; i <= 10000000; i++ {
		fmt.Println(i)
	}

	elapsed := time.Since(start).Seconds()
	fmt.Printf("Tiempo transcurrido: %.9fs\n", elapsed)
}
