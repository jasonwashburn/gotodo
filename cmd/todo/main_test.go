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
	binName  = "gotodo"
	fileName = "test-todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	os.Setenv("TODO_FILENAME", fileName)

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "test task Number 1"

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

		expected := fmt.Sprintf("  1: %s\n", task)

		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead.", expected, string(out))
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		task2 := "test task Number 2"
		taskCmd := exec.Command(cmdPath, "-task", task2)
		if err := taskCmd.Run(); err != nil {
			t.Fatal(err)
		}
		completeCmd := exec.Command(cmdPath, "-complete", "1")
		if err := completeCmd.Run(); err != nil {
			t.Fatal(err)
		}

		listCommand := exec.Command(cmdPath, "-list")

		out, err := listCommand.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := "X 1: test task Number 1\n  2: test task Number 2\n"

		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead.", expected, string(out))
		}
	})
}
