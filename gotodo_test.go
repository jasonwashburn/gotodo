package gotodo_test

import (
//    "io/ioutil"
//    "os"
    "testing"

    "github.com/jasonwashburn/gotodo"
)

func TestAdd(t *testing.T) {
    l := gotodo.List{}

    taskName := "New Task"
    l.Add(taskName)

    if l[0].Task != taskName {
        t.Errorf("Expected %q, but got %q instead.", taskName, l[0].Task)
    }
}

func TestComplete(t *testing.T) {
    l := gotodo.List{}

    taskName := "New Task"
    l.Add(taskName)

    if l[0].Task != taskName {
        t.Errorf("Expected %q, but got %q instead.", taskName, l[0].Task)
    }

    if l[0].Done {
        t.Errorf("New task should not be completed.")
    }

    l.Complete(1)

    if !l[0].Done {
        t.Errorf("New task should be completed.")
    }

}
