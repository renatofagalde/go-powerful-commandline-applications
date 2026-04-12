package main

import (
	"bytes"
	"testing"
)

func TestCoutingWords(t *testing.T) {
	bufferString := bytes.NewBufferString("word1 word2 word3\n")

	exp := 3

	result := count(bufferString)

	if result != exp {
		t.Errorf("words count failed, expected %d, got %d", exp, result)
	}

}
