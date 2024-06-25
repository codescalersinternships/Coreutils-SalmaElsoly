package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"log"
)
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}


func main() {
	var numberOfLines int
	flag.IntVar(&numberOfLines,"n", 10, "numberOfLines")
	flag.Parse()
	
	readFile, err := os.Open(flag.Args()[0])
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
