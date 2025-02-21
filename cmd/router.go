package cmd

func ParseCommand(args []string) {

	if len(args) == 1 {
		ListTasks([]string{})
		return
	}
	GetCommand(args[1], args[2:])
}
