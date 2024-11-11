package features

import (
	"bufio"
	"fmt"
	"os"
)

// Color constants
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Bold    = "\033[1m"
)

func ViewAllNotes(collection string) {
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

	fmt.Println(Green + "\nNotes:" + Reset)
	for i, note := range notes {
		fmt.Printf("%03d - %s\n", i+1, note)
	}
}
