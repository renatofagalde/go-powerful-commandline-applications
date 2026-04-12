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
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}

func count(stdin io.Reader, countLines bool) int {

	scanner := bufio.NewScanner(stdin)
	if !countLines {

		scanner.Split(bufio.ScanWords)
	}

	wordCounter := 0
	for scanner.Scan() {
		wordCounter++
	}
	return wordCounter
}
