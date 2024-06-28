package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for {
		if len(args) > 0 {
			fmt.Println(args[0])
		} else {
			fmt.Println("y")
		}
	}

}
