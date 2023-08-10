package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wordCount()
}
func wordCount() {
	for _, fileName := range os.Args[1:] {
		var lc, cc, wc int
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text()
			lc++
			cc += len(s)
			wc += len(strings.Fields(s))

		}
		fmt.Printf("%5d %5d %5d %s\n", lc, cc, wc, fileName)
		file.Close()
	}
}
