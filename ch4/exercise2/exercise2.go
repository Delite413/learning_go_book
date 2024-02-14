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

	for _, number := range intSlice {
		switch {
		case number%2 == 0 && number%3 == 0:
			fmt.Println("Six")
		case number%2 == 0:
			fmt.Println("Two")
		case number%3 == 0:
			fmt.Println("Three")
		default:
			fmt.Println("Never Mind")
		}
	}
}
