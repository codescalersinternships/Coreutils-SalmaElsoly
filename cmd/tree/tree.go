package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
		lisit(1, level, directoriesOnly, flag.Args()[0])
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
