package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var (
		level           int
		directoriesOnly bool
	)
	flag.BoolVar(&directoriesOnly, "d", false, "directoriesOnly")
	flag.IntVar(&level, "L", 1, "level")

	flag.Parse()
	if len(flag.Args()) > 0 {
		getNestedDirectoriesAndFiles(1, level, directoriesOnly, flag.Args()[0])
	} else {
		curr, err := os.Getwd()
		check(err)
		getNestedDirectoriesAndFiles(1, level, directoriesOnly, curr)
	}

}

func getNestedDirectoriesAndFiles(currLevel int, maxLevel int, directoriesOnly bool, path string) {
	if currLevel > maxLevel {
		return
	}
	content, err := os.ReadDir(path)
	check(err)
	for _, entry := range content {
		if entry.IsDir() {
			fmt.Print(strings.Repeat("   ", currLevel))
			fmt.Println("|__", entry.Name())
			getNestedDirectoriesAndFiles((currLevel + 1), maxLevel, directoriesOnly, filepath.Join(path, entry.Name()))
		} else {
			if !directoriesOnly {
				fmt.Print(strings.Repeat("   ", currLevel))
				fmt.Println("|__", entry.Name())
			}
		}
	}
}
