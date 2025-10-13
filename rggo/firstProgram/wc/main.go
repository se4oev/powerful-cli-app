package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	countBytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *countBytes))
}

func count(r io.Reader, countLines, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	counter := 0
	scanner.Bytes()
	for scanner.Scan() {
		if countBytes {
			counter += len(scanner.Bytes())
			continue
		}
		counter++
	}
	return counter
}
