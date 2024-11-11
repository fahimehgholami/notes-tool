package main

import (
	"bufio"
	"fmt"
	"notestool/features"
	"os"
	"strings"
)

// Color constants
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Cyan    = "\033[36m"
	Magenta = "\033[35m"
	Bold    = "\033[1m"
)

func PrintHelp() {
	fmt.Println(Bold + "Usage:\n\n" + Reset + "./notestool [COLLECTION_NAME]\n\n" + Bold + "[COLLECTION_NAME]" + Reset + " = Mandatory argument for use. Must be a non-existent or a non-directory plain text file that is available.\n\nExample usage:\n./notestool shoppinglist\n\n" + Bold + "Notes tool - A command-line tool to manage notes." + Reset)
}

func main() {
	// Save the name of the file to be used as database into var collection string
	var collection string

	// Handle cases where we may want to print Help/Usage and terminate the program
	if len(os.Args) != 2 {
		fmt.Println(Red + "Invalid number of arguments." + Reset)
		PrintHelp()
		return
	} else {
		collection = os.Args[1]
		if collection == "help" {
			PrintHelp()
			return
		}
	}

	// We have collected a single parameter collection, but lets see if its available and usable as a Notes tool database file
	// BasicEdgeCaseHandling will handle some common edge cases like parameter pointing to directory, binary file, file over 999 lines
	if !features.BasicEdgeCaseHandling(collection, os.Args[0]) {
		// BasicEdgeCaseHandling prints its own error messages if any, so we just terminate main function/program in case it returns false
		return
	}

	// Happy path starts from here
	fmt.Printf("Welcome to the notes tool for collection: %s%s%s\n", Magenta, collection, Reset)
	for {
		fmt.Println("\n" + Cyan + "Select operation:" + Reset)
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			features.ViewAllNotes(collection)
		case "2":
			features.AddNoteToCollection(collection)
		case "3":
			features.DeleteNoteFromCollection(collection)
		case "4":
			fmt.Println(Magenta + "Exiting the notes tool." + Reset)
			return
		default:
			fmt.Println(Red + "Invalid option, please select 1-4." + Reset)
		}
	}
}
