package main

import (
	"fmt"
	"io"
	"os"
)

const FILE_NAME = "Textfile.txt"

func CheckErrNil(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadTextFile() {
	fmt.Println("Reading the Text file")
	dataBytes, err := os.ReadFile(FILE_NAME)
	CheckErrNil(err)

	fmt.Println("File Content: ", string(dataBytes))
}

func main() {

	fmt.Println("___________FileOperations__________")

	// create the file
	file, err := os.Create(FILE_NAME)
	CheckErrNil(err)

	// close the file at the end of the program
	defer file.Close()

	// Content to write in the file
	content := "This is my Content"

	// write to the file
	length, err := io.WriteString(file, content)
	CheckErrNil(err)
	fmt.Println("Length of the file content", length)

	// Read the text file
	ReadTextFile()
}
