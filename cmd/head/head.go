package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"log"
	"errors"
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
	if len(flag.Args())< 1{
		errNoArg:= errors.New("No file specified")
		log.Fatal(errNoArg)
	} 
	readFile, err := os.Open(flag.Args()[0])
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	for i := 0; i < numberOfLines; i++ {
		if !fileScanner.Scan() {
			break
		}
		fmt.Println(fileScanner.Text())
	}

}
