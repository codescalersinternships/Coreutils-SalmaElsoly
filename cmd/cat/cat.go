package main

import (
	"bufio"
	"errors"
	"flag"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var numberedLines bool
	flag.BoolVar(&numberedLines, "n", false, "numberedFiles")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1{
		errNoArg:= errors.New("No file specified")
		log.Fatal(errNoArg)
	}
	fileToRead, err := os.Open(args[0])
	check(err)
	defer fileToRead.Close()
	
	fileScanner := bufio.NewScanner(fileToRead)
	count := 0
	for fileScanner.Scan() {
		if numberedLines {
			count++
			os.Stdout.WriteString(strconv.Itoa(count) + " " + fileScanner.Text()+"\n")
		} else {
			os.Stdout.WriteString(fileScanner.Text()+"\n")
		}

	}
	
}
