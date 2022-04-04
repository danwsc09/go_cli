package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Define a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")

	// Defile a boolean flag -b to count the number of bytes
	bytes := flag.Bool("b", false, "Count bytes")

	// Parse flags provided by user
	flag.Parse()

	// Calling the count function to count the number of words (or lines)
	// received from the Standard Input and printting it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// Scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// if the count lines flag is not set, we want to count words, so we define
	// the scanner split type to words (default is split by lines)
	if !countLines {
		if countBytes {
			scanner.Split(bufio.ScanBytes)
		} else {
			scanner.Split(bufio.ScanWords)
		}
	}

	// Define a counter
	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
