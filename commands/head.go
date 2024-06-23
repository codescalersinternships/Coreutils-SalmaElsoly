package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)


func Head(){
	headCmd:=flag.NewFlagSet("head",flag.ExitOnError)
	numberOfLines:=headCmd.Int("n",10,"numberOfLines")
	headCmd.Parse(os.Args[2:])
	path:=headCmd.Args()[0]
	absPath,err:= filepath.Abs(path)
	check(err)

	readFile,err:=os.Open(absPath)
	check(err)

	fileScanner := bufio.NewScanner(readFile)

	for i:=0; i < *numberOfLines;i++{
		if(fileScanner.Scan()){
			fmt.Println(fileScanner.Text())
		}else{
			break
		}
	
	}

}
