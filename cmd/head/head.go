package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)
func check(e error) {
	if e != nil {
		log.fatal(e)
	}
}

var(
	numberOfLines int
)

func main() {
	headCmd := flag.NewFlagSet("head", flag.ExitOnError)
	headCmd.IntVar(&numberOfLines,"n", 10, "numberOfLines")
	headCmd.Parse(os.Args[1:])
	path := headCmd.Args()[0]
	absPath, err := filepath.Abs(path)
	check(err)

	readFile, err := os.Open(absPath)
	check(err)

	fileScanner := bufio.NewScanner(readFile)

	for i := 0; i < numberOfLines; i++ {
		if fileScanner.Scan() {
			fmt.Println(fileScanner.Text())
		} else {
			break
		}

	}

}
