package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(count(os.Stdin))
}

func count(stdin io.Reader) int {
	scanner := bufio.NewScanner(stdin)
	scanner.Split(bufio.ScanWords)

	wordCounter := 0
	for scanner.Scan() {
		wordCounter++
	}
	return wordCounter
}
