package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var newLineTrailing bool
	flag.BoolVar(&newLineTrailing, "n", false, " omit the trailing newline")
	flag.Parse()
	args := flag.Args()

	if newLineTrailing {
		fmt.Print(strings.Join(args, " "))
		return
	}
	for i := 0; i < len(args); i++ {

		fmt.Println(args[i])

	}
}
