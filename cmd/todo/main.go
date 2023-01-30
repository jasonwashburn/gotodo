package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jasonwashburn/gotodo"
)

func main() {
	const todoFileName = ".todo.json"
	l := &gotodo.List{}

	task := flag.String("task", "", "Task to be included in the TODO list")
	list := flag.Bool("list", false, "List all the tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(l)
	case *task != "":
		l.Add(*task)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
