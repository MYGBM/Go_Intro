package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	sizeOfFile()
}
func sizeOfFile() {
	for _, fileName := range os.Args[1:] {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println("This file has ", len(data), "bytes") // since data needs to be used outside of the if statement, a short declaration with an if statement isn't used
		file.Close()

	}
}
