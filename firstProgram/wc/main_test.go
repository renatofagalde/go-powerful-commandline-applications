package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	bufferString := bytes.NewBufferString("word1 word2 word3\n")

	exp := 3

	result := count(bufferString, false)

	if result != exp {
		t.Errorf("words count failed, expected %d, got %d", exp, result)
	}

}

func TestCountLines(t *testing.T) {
	bufferString := bytes.NewBufferString("word1 word2 word3\nword1 word2 word3\nword1 word2 word3\n")

	exp := 3

	result := count(bufferString, true)

	if result != exp {
		t.Errorf("lines count failed, expected %d, got %d", exp, result)
	}
}
