package ideas

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var File *os.File

// append idea to file
func Add(idea string) {
	_, err := File.WriteString(idea + "\n")
	checkError(err)
	fmt.Println("Added idea")
}

func Remove(index int) {
	scanner := bufio.NewScanner(File)

	// get home dir
	homeDir, err := os.UserHomeDir()
	checkError(err)

	// create temp file
	temp, err := os.CreateTemp(homeDir, "godeas")
	checkError(err)

	// scan file line by line
	i := 1
	for scanner.Scan() {
		if i == index { // the line we want to remove
			fmt.Print("Are you sure you want to remove \"" + scanner.Text() + "\"? [y/N]: ")

			// get user input
			reader := bufio.NewReader(os.Stdin)
			char, _, err := reader.ReadRune()
			checkError(err)

			// if yes then remove
			if char == 'y' || char == 'Y' {
				fmt.Println("Removed: " + "\"" + scanner.Text() + "\"")
			}
		} else { // all other lines
			// write current line to file
			_, err = temp.WriteString(scanner.Text() + "\n")
			checkError(err)
		}

		// update index
		i++
	}

	// rename temp file to actual file
	err = os.Rename(temp.Name(), File.Name())
	checkError(err)
}

func List() {
	scanner := bufio.NewScanner(File)
	// scan the file line by line
	i := 1
	for scanner.Scan() {
		fmt.Println(i, scanner.Text())
		i++
	}
}

func Clear() {
	fmt.Print("Are you sure you want to clear all ideas? [y/N]: ")

	// get user input
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	checkError(err)

	// check user input
	if char == 'y' || char == 'Y' {
		os.Create(File.Name())
		fmt.Println("Cleared ideas")
	} else {
		fmt.Println("Didn't clear ideas")
	}
}

// check error helper function
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
