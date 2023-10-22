package ideas

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var path string
var ideas []string

func SetPath(p string) {
	path = p
}

func Read() {
	// open path
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	checkError(err)
	defer file.Close()

	// scan the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ideas = append(ideas, scanner.Text())
	}
}

func Add(idea string) {
	// append idea to ideas array
	ideas = append(ideas, idea)
	fmt.Println("Added idea:", idea)
}

func Remove(index int) {
	fmt.Printf("Are you sure you want to remove idea: \"%v\"? [y/N]: ", ideas[index-1])

	// get user input
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	checkError(err)

	// check user input
	if char == 'y' || char == 'Y' {
		// if the length of the ideas array is 1 then just clear it because if we remove it we get an out of range exception
		if len(ideas) > 1 {
			// remove element from ideas array
			ideas = append(ideas[:index-1], ideas[index:]...)
		} else {
			ideas = nil
		}
	} else {
		fmt.Println("Didn't remove idea")
	}
}

func List() {
	// list all ideas and print index
	for i, v := range ideas {
		index_num := fmt.Sprintf("[%d]", i+1)
		fmt.Println(index_num, v)
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
		ideas = nil
	} else {
		fmt.Println("Didn't clear ideas")
	}
}

func Backup() {
	input, err := os.ReadFile(path)
	checkError(err)

	err = os.WriteFile(path+".backup", input, 0600)
	checkError(err)
}

func Store() {
	// remove old file before writing to new one
	err := os.Remove(path)
	checkError(err)

	// create new file
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	checkError(err)
	defer file.Close()

	// write ideas array to file
	for _, v := range ideas {
		file.WriteString(v + "\n")
	}
}

// check error helper function
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
