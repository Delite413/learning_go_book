package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "testFile.txt"
	bytes, err := fileLen(fileName)
	if err != nil {
		fmt.Println("There was an error", err)
		os.Exit(1)
	}

	fmt.Println(bytes)
}

func fileLen(fileName string) (int64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return 0, err
	}

	fileSize := fileInfo.Size()
	return fileSize, nil
}
