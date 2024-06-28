package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		for _, e := range os.Environ() {
			fmt.Println(e)
		}
		return
	}
	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], "=") {
			pair := strings.Split(args[i], "=")
			if len(pair) == 2{
				os.Setenv(pair[0], pair[1])
				fmt.Println(pair[0], "=", pair[1])
			}
		} else {
			fmt.Println(os.Getenv(args[i]))
			str := os.Getenv(args[i])
			if strings.Contains(str, ";") {
				values := strings.Split(str, ";")
				for j := range values {
					fmt.Println(args[i], "=", values[j])
				}
			} else {
				fmt.Println(args[i], "=", str)
			}
		}

	}
}
