# Notes Tool

This tool manages short single-line notes.

Using this tool, you can create a collection of notes.

Operations available for a given notes collection:

1. Show notes.
2. Add a note.
3. Delete a note.

### How Notes data is stored

* Notes are saved locally into plain text files within the same folder where the tool is run.
* Collections are stored within a file which have the name of the collection.
* Each row within a collection plain text file represents a single note within a collection.

### Basic Usage

Usage:

    ./notestool {COLLECTION_NAME}

A single argument ```{COLLECTION_NAME}``` is mandatory as each note will be part of and saved into a collection.

Eg.

    ./notestool shoppinglist

... would open Notes tool ready to edit the notes within the collection "shoppinglist".

Operation ```2. Add a note.``` will allow entering ```"banana"``` on the terminal.

Operation ```1. Show notes.``` will list ```"001 - banana"```.

Operation ```3. Delete a note.``` will allow entering ```"1"``` to remove note ```"banana"```.

Errors are displayed in case operation is not currently available or fails, eg. ```"No notes available."``` if collection is still empty of notes.

### Limitations

* Since the application uses plain text files to store the notes, the collection parameter given can not have the same name as an existing non-collection file or a directory.
* The program will print some errors if it detects it can not work with a given collection parameter due to this and quits.
* Program maximally handles 999 notes per collection and rejects collection parameter pointing to files which have more than 999 notes/lines.
* Hand holding is otherwise limited - careful usage according to common sense is advised.

# Credits

Rustholkarhu, Gholami, Pandit

## About

Coded at Kood/Sisu Sprint IV 2024 as the collaborative Group Task 2