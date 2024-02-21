package main

import (
	"errors"
	"fmt"
)

var ErrInvalidQuote = errors.New("invalid quote")

func main() {
	err := createPolicy(11)
	if err != nil {
		if errors.Is(err, ErrInvalidQuote) {
			fmt.Println("threw custom error")
		} else {
			fmt.Println("other random error")
		}
	}

	fmt.Println("ran fine")
}

func createPolicy(referenceNumber int) error {
	if referenceNumber%2 == 0 {
		return nil
	}

	if referenceNumber%2 == 1 {
		return ErrInvalidQuote
	}

	return nil
}
