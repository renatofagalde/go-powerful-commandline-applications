package todo_test

import (
	todo "bootstrap"
	"testing"
)

func TestAdd(t *testing.T) {
	list := todo.List{}

	var taskName string = "New Task"
	list.Add(taskName)

	if list[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, list[0].Task)
	}
}
