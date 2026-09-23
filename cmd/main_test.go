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
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {

	fmt.Println("Creazione dello strumento...")

	if runtime.GOOS == "windows" {
		binName += ".exe "
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Impossibile compilare lo strumento %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Esecuzione dei test...")
	risultato := m.Run()

	fmt.Println("Pulizia in corso...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(risultato)
}
func TestTodoCLI(t *testing.T) {
	task := "test task number 1"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-task", task)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := task + "\n"

		if expected != string(out) {
			t.Errorf("Expected %q, got %q", expected, string(out))
		}
	})
}
