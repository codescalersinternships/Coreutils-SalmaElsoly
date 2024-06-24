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
		default:
			fmt.Println("Enter valid command")

	}
}