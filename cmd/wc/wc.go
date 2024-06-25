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
	lines bool
	words bool
	characters bool
)
func main(){
	wcCmd := flag.NewFlagSet("wc", flag.ExitOnError)
	lines := wcCmd.BoolVar(&lines,"l", false, "lines")
	words := wcCmd.BoolVar(&words,"w", false, "words")
	characters := wcCmd.BoolVar(&characters,"c", false, "characters")

	wcCmd.Parse(os.Args[1:])
	path := wcCmd.Args()[0]

	absPath, err := filepath.Abs(path)
	check(err)

	if lines {
		readFileLines, err := os.Open(absPath)
		check(err)
		lineScanner := bufio.NewScanner(readFileLines)
		lineScanner.Split(bufio.ScanLines)
		linesC := 0
		for lineScanner.Scan() {
			linesC++
		}
		fmt.Println("lines count: ", linesC)
	}
	if words {
		readFileWords, err := os.Open(absPath)
		check(err)
		wordScanner := bufio.NewScanner(readFileWords)
		wordScanner.Split(bufio.ScanWords)
		wordsC := 0
		charsC := 0
		for wordScanner.Scan() {
			wordsC++
			charsC += len(wordScanner.Text())
		}
		fmt.Println("words count:", wordsC)
		if characters {
			fmt.Println("chars count:", charsC)
		}
	}
}
