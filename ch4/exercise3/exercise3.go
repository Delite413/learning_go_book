package main

import "fmt"

func main() {
	total := 0
	for i, total := 0, 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}

	// This will be shadowed
	fmt.Println(total)
}
