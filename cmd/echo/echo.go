package main

import (
	"flag"
	"os"
)

func main() {
	var newLineTrailing bool
	flag.BoolVar(&newLineTrailing, "n", false, " omit the trailing newline")
	flag.Parse()
	args := flag.Args()

	for i := 0; i < len(args); i++ {
		if newLineTrailing {
			os.Stdout.WriteString(args[i]+" ")
		} else {
			os.Stdout.WriteString(args[i]+"\n")
		}
	}
}
