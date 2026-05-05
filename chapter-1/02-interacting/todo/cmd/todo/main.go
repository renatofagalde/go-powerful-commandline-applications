package main

import (
	todo "bootstrap"
	"fmt"
	"os"
	"strings"
)

const TODO_FILE_NAME string = ".todo.json"

func main() {
	var list *todo.List = &todo.List{}
	if err := list.Get(TODO_FILE_NAME); err != nil && !os.IsNotExist(err) {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {

	case len(os.Args) == 1:
		for _, item := range *list {
			fmt.Println(item.Task)
		}

	default:
		item := strings.Join(os.Args[1:], " ")
		fmt.Println(item)

		list.Add(item)

		if err := list.Save(TODO_FILE_NAME); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
