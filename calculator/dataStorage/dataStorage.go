package dataStorage

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func SaveCalculation(userInputs string) {
	content := userInputs
	filePath := "./calculations.txt"

	/*
		isFileExist := checkFileExists(filePath)

		if isFileExist {
			fmt.Println("file exist")
		} else {

			fmt.Println("file not exists")
		}
	*/

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close()
	// defer readFile(filePath)
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:\n", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}
