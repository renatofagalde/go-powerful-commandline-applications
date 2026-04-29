package main_test

import (
	"fmt"
	"runtime"
	"testing"
)

var (
	binName  string = "todo"
	fileName string = "todo_test.go"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

}
