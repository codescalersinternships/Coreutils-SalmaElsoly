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
		fmt.Println(e)
		fmt.Println("error occured")
		panic(e)

	}
}

func main() {
	tailCmd := flag.NewFlagSet("tail", flag.ExitOnError)
	numberOfLines := tailCmd.Int("n", 10, "numberOfLines")
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

	for i := 1; i <= *numberOfLines; i++ {
		if i < len(linesOfFile) {
			fmt.Println(linesOfFile[len(linesOfFile)-i])
		}

	}

}
