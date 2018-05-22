package main

import (
	"os/user"
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	current_user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Lepto programming language!\n", current_user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
