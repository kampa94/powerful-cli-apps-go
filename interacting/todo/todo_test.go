package todo_test

import (
	"os"
	"testing"
	"todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected task name %q, got %q", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	l.Add("Task to complete")
	err := l.Complete(0)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !l[0].Done {
		t.Errorf("Expected task to be marked as done")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}
	l.Add("Task to delete")
	err := l.Delete(0)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(l) != 0 {
		t.Errorf("Expected list to be empty after deletion, got length %d", len(l))
	}
}

func TestSaveAndGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}
	taskName := "Task to save"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected task name %q, got %q", taskName, l1[0].Task)
	}
	tf, err := os.CreateTemp("", "todo_test_*.json")
	if err != nil {
		t.Fatalf("Expected no error creating temp file, got %v", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Expected no error saving file, got %v", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Expected no error getting file, got %v", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Expected task name %q, got %q", l1[0].Task, l2[0].Task)
	}
}
