package main

import (
	"coreutils-go/commands"
	"fmt"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

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
		default:
			fmt.Println("Enter valid command")

	}
}