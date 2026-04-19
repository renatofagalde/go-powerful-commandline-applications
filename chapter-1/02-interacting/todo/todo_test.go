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

func TestComplete(t *testing.T) {
	list := todo.List{}

	var taskName string = "New Task"
	list.Add(taskName)

	if list[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, list[0].Task)
	}
	if list[0].Done {
		t.Errorf("New task shouldn't be completed")
	}

	_ = list.Complete(1)
	if !list[0].Done {
		t.Errorf("New task shouldn't be completed")
	}
}
