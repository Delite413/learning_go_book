package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intSlice := make([]int, 0, 100)
	for range 100 {
		randomInt := rand.Intn(100)
		intSlice = append(intSlice, randomInt)
	}

	fmt.Println(intSlice)
}
