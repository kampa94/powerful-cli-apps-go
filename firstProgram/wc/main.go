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
	bytes := flag.Bool("b", false, "Count bytes")
	chars := flag.Bool("c", false, "Count chars")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *chars, *bytes))
}

func count(r io.Reader, countLines bool, countRunes bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	switch {
	case countLines:
		scanner.Split(bufio.ScanLines)
	case countRunes:
		scanner.Split(bufio.ScanRunes)
	case countBytes:
		scanner.Split(bufio.ScanBytes)
	default:
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
