package cmd

import (
	"fmt"

	"github.com/javiersrf/todo-cli/internal"
)

func AddTodo(args []string) {
	task := args[0]
	fmt.Printf("Adding task '%s' to the list\n", task)
	db, err := internal.InitDB()
	if err != nil {
		panic(err)
	}
	internal.AddTask(db, task)

}
