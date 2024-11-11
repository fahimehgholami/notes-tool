package features

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddNoteToCollection(collection string) {
	fmt.Println(Blue + "Enter the note text:" + Reset)

	reader := bufio.NewReader(os.Stdin)
	noteText, _ := reader.ReadString('\n')
	noteText = strings.TrimSpace(noteText)

	file, err := os.OpenFile(collection, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(Red+"Error opening the collection file:"+Reset, err)
		return
	}
	defer file.Close()

	// Edge Case: Quickly read how many notes already exist and prevent adding a note if 999 (or more) notes already exist
	scanner := bufio.NewScanner(file)
	check_notes := make([]string, 0)
	for scanner.Scan() {
		check_notes = append(check_notes, scanner.Text())
	}
	if len(check_notes) >= 999 {
		fmt.Println(Red + "Collection has 999 or more notes. Please start a new Collection or Delete notes!" + Reset)
		return
	}

	// This will physically write the note to the file
	_, err = file.WriteString(noteText + "\n")
	if err != nil {
		fmt.Println(Red+"Error writing to the file:"+Reset, err)
	}
	//fmt.Println("Note added successfully.")
}
