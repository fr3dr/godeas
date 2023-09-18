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
	flag.Parse()

	// get home dir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// read idea file
	file, err := os.OpenFile(homeDir+"/ideas.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	ideas.File = file

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
}
