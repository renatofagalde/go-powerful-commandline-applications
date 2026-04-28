package main

import (
	todo "bootstrap"
	"fmt"
	"os"
)

const TODO_FILE_NAME string = ".todo.json"

func main() {
	var list todo.List = todo.List{}

	if err := list.Get(TODO_FILE_NAME); err != nil {
		//errorString := fmt.Sprintf("Error %v to open the %s file", err, TODO_FILE_NAME)
		//fmt.Println(errorString)

		//Whe developing a command-line tool,
		//it's a good practice to use the standard error(STDERR)
		//output insetead of the standard output(STDOUT)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
