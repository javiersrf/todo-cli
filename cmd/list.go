package cmd

import (
	"fmt"

	"github.com/javiersrf/todo-cli/internal"
)

func ListTasks(args []string) {
	db, err := internal.InitDB()
	if err != nil {
		panic(err)
	}
	tasks, err := internal.GetTasks(db)
	if err != nil {
		panic(err)
	}

	println("--------------------------------- Todo-Cli -----------------------------------")

	fmt.Printf("Total tasks: %d\n", len(tasks))
	println("ID\t\t\t\t\tTask\t\t\t\tStatus")

	for _, task := range tasks {
		status := "[ ]"
		if task.Status == 1 {
			status = "[x]"
		}
		// println(task.ID, "\t\t\t", task.Task, "\t\t\t\t", status)
		fmt.Printf("%d\t\t\t\t\t%s\t\t\t\t%s", task.ID, task.Task, status)
		println()
	}
	println("----------------------------------END-----------------------------------------")

}
