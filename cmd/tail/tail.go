package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
var(
	numberOfLines int
)
func main() {
	tailCmd := flag.NewFlagSet("tail", flag.ExitOnError)
	tailCmd.IntVar(&numberOfLines,"n", 10, "numberOfLines")
	tailCmd.Parse(os.Args[1:])
	path := tailCmd.Args()[0]
	absPath, err := filepath.Abs(path)
	check(err)

	readFile, err := os.Open(absPath)
	check(err)

	fileScanner := bufio.NewScanner(readFile)

	var linesOfFile []string

	for fileScanner.Scan() {
		linesOfFile = append(linesOfFile, fileScanner.Text())
	}

	for i := 0; i < numberOfLines; i++ {
		if i < len(linesOfFile) {
			fmt.Println(linesOfFile[len(linesOfFile)-numberOfLines+i])
		}

	}

}
