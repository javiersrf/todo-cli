package cmd

import (
	"fmt"
	"strconv"

	"github.com/javiersrf/todo-cli/internal"
)

func MarkTask(args []string) {
	taskId := args[0]
	db, err := internal.InitDB()
	if err != nil {

		panic(err)
	}
	id, err := strconv.Atoi(taskId)
	if err != nil {
		panic(err)
	}
	task, err := internal.GetTask(db, id)
	if err != nil {
		println("Task not found")
		return
	}
	newStatus := 1
	if task.Status == 1 {
		newStatus = 0
	}
	internal.MarkTask(db, id, newStatus)
	task, err = internal.GetTask(db, id)
	if err != nil {
		panic(err)
	}
	finalStatus := "[ ]"
	if task.Status == 1 {
		finalStatus = "[x]"
	}
	fmt.Printf("Task %s  -> %s", taskId, finalStatus)
	println()

}
