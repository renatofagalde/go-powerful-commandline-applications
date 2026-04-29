package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Cannot remove the tool %s: %s\n", name, err)
		}
	}(fileName + ".test")

	os.Exit(run)
}

func TestTodoCLI(t *testing.T) {
	var task string = "Test task number 1"
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	if cmdPath == "" {
		t.Fatal("Cannot find the tool")
	}

	t.Run("Add New Task", func(t *testing.T) {
		exec.Command(cmdPath, strings.Split(task, " ")...).Run()

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
}
