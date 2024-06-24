package main

import (
	"coreutils-go/commands"
	"fmt"
	"os"
)

func main(){
	switch os.Args[1]{
		case "head":
			cmd.Head()
		case "tail":
			cmd.Tail()
		case "wc":
			cmd.WordCount()
		case "cat":
			cmd.Cat()
		case "echo":
			cmd.Echo()
		case "env":
			cmd.Env()
		case "yes":
			cmd.Yes()
		case "false":
			os.Exit(3)
		case "true":
			os.Exit(0)
		default:
			fmt.Println("Enter valid command")

	}
}