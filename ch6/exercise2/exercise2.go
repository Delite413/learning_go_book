package main

import "fmt"

func main() {
	stringCollection := []string{"Apple", "Banana", "Cucumber"}
	fmt.Println(stringCollection)
	UpdateSlice(stringCollection, "Function1")

	fmt.Println(stringCollection)
	GrowSlice(stringCollection, "Function2")
}

func UpdateSlice(stringCollection []string, inputString string) {
	stringCollection[len(stringCollection)-1] = inputString
	fmt.Println(stringCollection)
}

func GrowSlice(stringCollection []string, inputString string) {
	stringCollection = append(stringCollection, inputString)
	fmt.Println(stringCollection)
}
