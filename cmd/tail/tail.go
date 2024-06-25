package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var numberOfLines int
	flag.IntVar(&numberOfLines, "n", 10, "numberOfLines")
	flag.Parse()
	path := flag.Args()[0]
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
