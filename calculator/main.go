/*
A calculator application written in Golang
The main purpose of this application is to further my knowledge on Golang

Planning/Functionalities:
  - Basic operations of a calculator ( + - * / )
  - Get user input from command line
  - Store User's previous calculations in a separate file
  - TBA Functionality:
  - Possible features, isEven, isOdd, find root, fibonacci, squared
*/

package main

import (
	"bufio"
	menu "calculator/menu"
	"fmt"
	"os"
	"strings"
)

func main() {
	runApp()
}

func runApp() {
	fmt.Print("Calculator App CLI Menu\n\n")

	run := true
	reader := bufio.NewReader(os.Stdin)

	for run {
		fmt.Print("Select what you'd like to do:\n\n")
		fmt.Print("(C)alculations\n(Q)uit\n\n")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ToUpper(input)

		switch input {
		case "C":
			fmt.Print("Going to calculations menu\n\n")
			menu.CalcMenu()
		case "Q":
			fmt.Print("Quitting application\n")
			run = false
		default:
			panic("Invalid input")
		}
	}
}
