package features

import (
	"bufio"
	"fmt"
	"os"
)

// Returns TRUE if we can proceed with input parameter [collection string] and FALSE if problems detected
func BasicEdgeCaseHandling(collection, self_path string) bool {
	if collection == self_path[2:] {
		fmt.Println(Red + "Fatal Error: Inputted collection parameter is the executable itself.")
		return false
	}
	var notes []string
	file, err := os.Open(collection)
	if err != nil {
		if os.IsNotExist(err) {
			// No file with collection string is a green flag for availability for notes collection
			return true
		}
		// Other problems with opening collection will be lumped into an unkown error
		fmt.Println(Red + "Fatal Error: Inputted collection parameter caused unknown file opening/analysation error." + Reset)
		return false
	}
	defer file.Close()

	// Now lets check we did not open a directory with os.Stat which I guess is similar to Unix stat command
	info, err := os.Stat(collection)
	if err != nil {
		fmt.Println(Red + "Fatal Error: Inputted collection parameter caused unknown file opening/analysation error." + Reset)
		return false
	}

	if info.IsDir() {
		fmt.Println(Red + "Fatal Error: Inputted collection parameter points to a directory, collecting notes under a directory is not supported." + Reset)
		return false
	}

	// A non-directory file was good for opening, so now finally lets read it also
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	// Over 999 lines/notes in file is suspicious could be eg. a large binary file, thus lets invalidate this as a usable collection file
	if len(notes) > 999 {
		fmt.Println(Red + "Fatal Error: Inputted collection parameter points to a large file, collecting notes under this collection will not be available." + Reset)
		return false
	}
	return true
}
