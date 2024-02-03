package main

import (
	"flag"
	"log"
	"os"

	"github.com/fr3dr/godeas"
)

func main() {
	// get home dir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// generate flags
	add := flag.String("a", "", "add an idea")
	remove := flag.Int("r", 0, "remove an idea by line number")
	list := flag.Bool("l", false, "list ideas")
	clear := flag.Bool("c", false, "clear ideas")
	path := flag.String("p", homeDir + "/ideas.txt", "custom path to ideas file")
	flag.Parse()

	// set path to ideas file
	ideas.SetPath(*path)

	// read file and back it up
	ideas.Read()
	ideas.Backup()

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
