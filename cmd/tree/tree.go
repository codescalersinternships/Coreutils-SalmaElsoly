package main

import (
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
	level int
	directoriesOnly bool
)
func main() {
	treeCmd := flag.NewFlagSet("tree", flag.ExitOnError)
	treeCmd.BoolVar(,&directoriesOnly"d", false, "directoriesOnly")
	treeCmd.IntVar(&level,"L", 1, "level")

	treeCmd.Parse(os.Args[1:])
	if len(treeCmd.Args()) > 0 {
		lisit(1, level, directoriesOnly, treeCmd.Args()[0])
	} else {
		curr, err := os.Getwd()
		fmt.Println(curr)
		check(err)
		lisit(1, level, directoriesOnly, curr)
	}

}

func lisit(currLevel int, maxLevel int, directoriesOnly bool, path string) {
	content, err := os.ReadDir(path)
	fmt.Println(content)
	check(err)
	if currLevel > maxLevel {
		return
	}

	for _, entry := range content {
		if entry.IsDir() {
			for i := 1; i < currLevel; i++ {
				fmt.Print(" ")
			}
			fmt.Println("|__", entry.Name())
			lisit((currLevel + 1), maxLevel, directoriesOnly, filepath.Join(path, entry.Name()))
		} else {
			if !directoriesOnly {
				for i := 1; i < currLevel; i++ {
					fmt.Println("|")
					fmt.Print("---", entry.Name())
				}
			}
		}
	}
}
