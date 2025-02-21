package main

import (
	"os"

	"github.com/javiersrf/todo-cli/cmd"
)

func main() {
	cmd.ParseCommand(os.Args)

}
