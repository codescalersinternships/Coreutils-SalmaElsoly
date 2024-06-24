package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	for range 10 {
		if len(args) > 0 {
			fmt.Println(args[0])
		} else {
			fmt.Println("yes")
		}
		time.Sleep(1 * time.Second)
	}

}
