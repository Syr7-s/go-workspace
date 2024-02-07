package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go!"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	defer func(file *os.File) {
		err := file.Close()
		checkError(err)
	}(file)
	defer readFile("./fromString.txt")
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Printf("Text read from file : %s", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
