package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func Tail(){
	tailCmd:=flag.NewFlagSet("tail",flag.ExitOnError)
	numberOfLines:=tailCmd.Int("n",10,"numberOfLines")
	tailCmd.Parse(os.Args[2:])
	path:=tailCmd.Args()[0]
	absPath,err:= filepath.Abs(path)
	check(err)

	readFile,err:=os.Open(absPath)
	check(err)

	fileScanner := bufio.NewScanner(readFile)

	var linesOfFile [] string

	for fileScanner.Scan(){
		linesOfFile = append(linesOfFile,fileScanner.Text())
	}

	for i:=1 ; i <= *numberOfLines ; i++{
		if(i<len(linesOfFile)){
			fmt.Println(linesOfFile[len(linesOfFile)-i])
		}
		
	}

}