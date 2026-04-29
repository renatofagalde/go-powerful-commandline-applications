package main_test

import (
	"fmt"
	"os"
	"os/exec"
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

	build := exec.Command("go", "build", "-o", binName, fileName)

	if err := build.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot build the tool %s: %s\n", binName, err)
		os.Exit(1)
	}

	fmt.Println("Tool built successfully\nRunning tests...")
	run := m.Run()

	fmt.Println("Cleaning up...")
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Cannot remove the tool %s: %s\n", name, err)
		}
	}(binName)
	defer os.Remove(fileName + ".test")

	os.Exit(run)
}
