package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	echoCmd := flag.NewFlagSet("head", flag.ExitOnError)
	trailing := echoCmd.Bool("n", false, "newLine")
	echoCmd.Parse(os.Args[1:])
	args := echoCmd.Args()

	for i := 0; i < len(args); i++ {
		if *trailing {
			fmt.Print(args[i], " ")
		} else {
			fmt.Println(args[i])
		}
	}
}
