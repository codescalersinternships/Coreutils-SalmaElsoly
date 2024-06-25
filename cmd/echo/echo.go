package main

import (
	"flag"
	"fmt"
)

func main() {
	var trailing bool
	flag.BoolVar(&trailing, "n", false, "newLine")
	flag.Parse()
	args := flag.Args()

	for i := 0; i < len(args); i++ {
		if trailing {
			fmt.Print(args[i], " ")
		} else {
			fmt.Println(args[i])
		}
	}
}
