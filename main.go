package main

import (
	"os"

	"github.com/javiersrf/todo-cli/internal"
)

func main() {
	internal.ParseCommand(os.Args)

}
