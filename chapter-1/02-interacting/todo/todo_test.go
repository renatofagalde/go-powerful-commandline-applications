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

func TestDelete(t *testing.T) {
	var l todo.List = todo.List{}

	//slice
	//var tasks []string = []string{"New Task 1", "New Task 2", "New Task 3"}
	// se delcarar com o tamanho é um array
	var tasks [3]string = [3]string{"New Task 1", "New Task 2", "New Task 3"}

	for _, task := range tasks {
		l.Add(task)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead", tasks[0], l[0].Task)
	}

	_ = l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected 2 items, got %d instead", len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead", tasks[2], l[1].Task)
	}

}
