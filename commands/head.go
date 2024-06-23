package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	numberOfLines := flag.Int("n",10,"an int")
	printLines(*numberOfLines, flag.Args()[0])
}

func check(e error) {
    if e != nil {
		fmt.Println(e)
        panic(e)
    }
}

func printLines(numberOfLines int, path string){
	absPath,err:= filepath.Abs(path)
	check(err)

	readFile,err:=os.Open(absPath)
	check(err)

	fileScanner := bufio.NewScanner(readFile)

	for i:=0; i < numberOfLines;i++{
		fileScanner.Scan()
		fmt.Println(fileScanner.Text())
	}

}