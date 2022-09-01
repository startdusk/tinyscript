package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/startdusk/tinyscript/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf(`
Hello %s! This is the Monkey programing language!
Feel free to type in command
`,
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
