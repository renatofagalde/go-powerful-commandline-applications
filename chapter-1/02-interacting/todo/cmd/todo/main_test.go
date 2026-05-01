package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  string = "todo"
	fileName string = ".todo.json"
)

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		err := os.WriteFile(fileName, []byte("[]"), 0644)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Cannot create the file %s", fileName)
			os.Exit(1)
		}
	}

	build := exec.Command("go", "build", "-a", "-o", binName)

	if err := build.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot build the tool %s: %s\n", binName, err)
		os.Exit(1)
	}

	run := m.Run()

	_ = os.Remove(binName)
	_ = os.Remove(fileName)

	os.Exit(run) // os.Exit doesn't call defer functions
}

func TestTodoCLI(t *testing.T) {
	var task string = "Test task number 1"
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("Add New Task", func(t *testing.T) {

		if err := exec.Command(cmdPath, task).Run(); err != nil {
			t.Fatal(err)
		}
		command := exec.Command("cat", fileName)
		out, _ := command.CombinedOutput()
		fmt.Println(fmt.Sprintf("File content at save: %s", string(out)))

	})

	t.Run("ListTasks", func(t *testing.T) {

		command := exec.Command(cmdPath)
		command.Dir = dir

		out, err := command.Output()
		if err != nil {
			t.Fatal(err)
		}

		expected := "Test task number 1\n"
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})
}
