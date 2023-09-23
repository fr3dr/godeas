package main

import (
	"flag"
	"log"
	"os"

	"github.com/fr3dr/godeas"
)

func main() {
	// generate flags
	add := flag.String("a", "", "add an idea")
	remove := flag.Int("r", 0, "remove an idea by line number")
	list := flag.Bool("l", false, "list ideas")
	clear := flag.Bool("c", false, "clear ideas")
	path := flag.String("p", "", "custom path to ideas file")
	flag.Parse()

	// set path to ideas file
	if *path != "" {
		ideas.SetPath(*path)
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		ideas.SetPath(homeDir + "/ideas.txt")
	}

	// read file
	ideas.Read()

	// check for flags
	switch {
	case *add != "":
		ideas.Add(*add)
	case *remove > 0:
		ideas.Remove(*remove)
	case *list:
		ideas.List()
	case *clear:
		ideas.Clear()
	}

	// store idea array back to file
	ideas.Store()
}
