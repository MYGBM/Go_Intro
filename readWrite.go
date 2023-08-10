package main

import (
	"fmt"
	"io"
	"os"
)

func readwrite() {
	//loop through the given commands
	for _, fileName := range os.Args[1:] {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue //continue to next file
		}
		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprint(os.Stderr, err)
			continue // continue to next file
		}
		file.Close() // for each file we need to close it before moving on to next file

	}

}
func main() {
	readwrite()

}
