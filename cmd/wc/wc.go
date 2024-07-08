package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var lines, words, characters bool
	flag.BoolVar(&lines, "l", false, "lines")
	flag.BoolVar(&words, "w", false, "words")
	flag.BoolVar(&characters, "c", false, "characters")

	flag.Parse()
	if len(flag.Args()) < 1 {
		errNoArg := errors.New("No file specified")
		log.Fatal(errNoArg)
	}
	var linesC, wordsC, charsC int
	readFileLines, err := os.Open(flag.Args()[0])
	check(err)
	defer readFileLines.Close()
	lineScanner := bufio.NewScanner(readFileLines)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		linesC++
		wordsC += len(strings.Split(lineScanner.Text(), " "))
		charsC += len([]byte(lineScanner.Text()))
	}
	if !lines && !words && !characters {
		fmt.Println("lines count: ", linesC)
		fmt.Println("words count:", wordsC)
		fmt.Println("chars count:", charsC)
	} else {
		if lines {
			fmt.Println("lines count: ", linesC)
		}
		if words {
			fmt.Println("words count:", wordsC)
		}
		if characters {
			fmt.Println("chars count:", charsC)
		}
	}

}
