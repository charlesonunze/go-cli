package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	// Command Line Params
	path := flag.String("path", "myapp.log", "The path to the log to be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, CRITICAL and ERROR.")

	flag.Parse()

	// "Open" a file and return an os.File value
	file, err := os.Open(*path)
	handleError(err)

	// Cleanup after final execution, this function will run after the main function is done executing
	defer file.Close()

	// Create a reader
	reader := bufio.NewReader(file)

	// Infinite loop
	for {
		// Read the string from the beginning of the file till it encounters a new line ie, '\n'
		string, err := reader.ReadString('\n')

		// This will break the infinite loop, more like a base case in recursion
		if err != nil {
			break
		}

		// Check if string contains the log level passed in
		if strings.Contains(string, *level) {
			fmt.Println(string)
		}
	}
}
