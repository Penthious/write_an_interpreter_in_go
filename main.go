package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/penthious/writing_an_interpreter_in_go/repl"
)

func main() {
	u, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", u.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

}