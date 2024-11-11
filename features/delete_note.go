package features

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DeleteNoteFromCollection(collection string) {
	var notes []string
	file, err := os.Open(collection)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(Red + "No notes available." + Reset)
			return
		}
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n" + Blue + "Enter the number of note to remove or 0 to cancel:" + Reset)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// strconv.ParseInt converts inputted "number of note to remove" string to int64
	numeric, err := strconv.ParseInt(input, 10, 32)
	switch {
	case input == "0":
		fmt.Println(Green + "No note was removed. Delete operation was canceled by user." + Reset)
		return
	case err == nil && numeric > 0:
		if len(notes) > int(numeric-1) {
			if len(notes) == 1 {
				// Delete collection file, do not leave a lingering empty DB upon last note deletion
				WipeCollection(collection)
				return
			}
			combined_notes := append(notes[0:numeric-1], notes[numeric:len(notes)]...)
			RewriteNotesFile(combined_notes, collection)
		} else {
			fmt.Println(Red + "No note found at entered index." + Reset)
			DeleteNoteFromCollection(collection)
		}
	default:
		fmt.Println(Red + "Invalid input." + Reset)
		DeleteNoteFromCollection(collection)
	}
}

func WipeCollection(collection string) {
	err := os.Remove(collection)
	if err != nil {
		fmt.Println(Red+"Error wiping the collection file: "+Reset, err)
	}
}

func RewriteNotesFile(noteText []string, collection string) {
	if err := os.Truncate(collection, 0); err != nil {
		fmt.Println(Red+"Error rewriting the collection file: "+Reset, err)
	}

	file, err := os.OpenFile(collection, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(Red+"Error opening the collection file: "+Reset, err)
		return
	}
	defer file.Close()

	for _, note := range noteText {
		_, err = file.WriteString(note + "\n")
		if err != nil {
			fmt.Println(Red+"Error writing to the file: "+Reset, err)
			return
		}
	}
}
