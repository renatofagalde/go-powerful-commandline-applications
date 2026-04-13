package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")

	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(stdin io.Reader, countLines bool, bytes bool) int {

	if bytes {
		data, _ := ioutil.ReadAll(stdin)
		return len(data)
	}

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
