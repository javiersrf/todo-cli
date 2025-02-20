package internal

import (
	"fmt"

	"github.com/javiersrf/todo-cli/cmd"
)

func ParseCommand(args []string) {
	fmt.Println(args)

	if len(args) == 1 {
		cmd.ShowWelcomeMessage()
		return
	}

	fmt.Println("Usage: todo-cli init")
	fmt.Println("Use 'init' to get started with Tasky.")
}
