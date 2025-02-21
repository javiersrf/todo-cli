package cmd

import "fmt"

type Command struct {
	Name    string
	Args    []string
	Help    string
	Execute func(args []string)
}

var commandList = []Command{
	{Name: "", Args: []string{}, Help: "List all the tasks", Execute: ListTasks},
	{Name: "add", Args: []string{"task"}, Help: "Add a new task to the list", Execute: AddTodo},
	{Name: "list", Args: []string{}, Help: "List all the tasks", Execute: ListTasks},
	{Name: "done", Args: []string{"task"}, Help: "Mark task as done", Execute: MarkTask},
	{Name: "remove", Args: []string{"task"}},
}

func GetCommand(command string, args []string) {
	for _, c := range commandList {
		if c.Name == command {

			if len(c.Args) != len(args) {
				ShowCommandUsage(c)
				return
			}
			if args[0] == "--help" {
				ShowHelp(c)
				return

			}
			c.Execute(args)
			return
		}
	}
}

func ShowCommandUsage(c Command) {
	fmt.Printf("Usage: todo-cli %s", c.Name)
	for _, arg := range c.Args {
		fmt.Printf(" <%s>", arg)
	}
	fmt.Println()
	fmt.Println("Type --help for more usages")
}

func ShowHelp(c Command) {
	fmt.Printf(c.Help)
	fmt.Println()
	fmt.Printf("Usage: todo-cli %s", c.Name)
	for _, arg := range c.Args {
		fmt.Printf(" <%s>", arg)
	}

	fmt.Println()
}
