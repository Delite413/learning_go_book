package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कर", "こんにちは", "Привіт"}
	slice1 := greetings[:2]
	slice2 := greetings[1:3]
	slice3 := greetings[3:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
