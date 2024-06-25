package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)
func check(e error) {
	if e != nil {
		log.fatal(e)
	}
}
var(
	numberedLines bool
)
func main() {
	catCmd := flag.NewFlagSet("wc", flag.ExitOnError)
	catCmd.BoolVar(&numberedLines,"n", false, "numberedFiles")

	catCmd.Parse(os.Args[1:])
	args := catCmd.Args()
	if len(args) > 1 {
		fileToWrite, errWrite := os.OpenFile(args[len(args)-1], os.O_CREATE, 0666)
		check(errWrite)
		if len(args) == 2 {
			data, err := os.ReadFile(args[0])
			check(err)
			_, errWrite := fileToWrite.WriteString(string(data))
			check(errWrite)
		} else {
			for i := 0; i < len(args)-1; i++ {
				data, err := os.ReadFile(args[i])
				check(err)
				_, errWrite := fileToWrite.WriteString(string(data))
				check(errWrite)
			}
		}
		defer fileToWrite.Close()
		readFile, _ := os.Open(args[len(args)-1])
		fileScanner := bufio.NewScanner(readFile)
		count := 0
		for fileScanner.Scan() {
			if numberedLines {
				count++
				fmt.Println(count, " ", fileScanner.Text())
			} else {
				fmt.Println(fileScanner.Text())
			}

		}

	} else {
		fileToRead, err := os.Open(args[0])
		check(err)
		fileScanner := bufio.NewScanner(fileToRead)
		count := 0
		for fileScanner.Scan() {
			if *numberedLines {
				count++
				fmt.Println(count, " ", fileScanner.Text())
			} else {
				fmt.Println(fileScanner.Text())
			}

		}
	}
}
