package menu

import (
	"bufio"
	calc "calculator/calculations"
	data "calculator/dataStorage"
	"fmt"
	"os"
	"strings"
)

func CalcMenu() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the first value: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	fmt.Print("Enter the second value: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	fmt.Print("Enter the operation ( + - * / ): ")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	fmt.Printf("\n1st Value: %v\n2nd Value: %v\nOperation: %v\n", input1, input2, operation)

	result := calc.Calculate(input1, input2, operation)
	fmt.Println("Result:", result)

	fmt.Print("\nWould you like to save calculation? (y/n): ")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	if answer == "y" {
		data.SaveCalculation(format(input1, input2, operation, result))
		fmt.Print("Calculation Saved!\n\n")
	} else if answer == "n" {
		fmt.Print("Returning to main menu.\n\n")
	} else {
		panic("Incorrect input")
	}
}

func format(input1, input2, operation string, result float64) string {
	var output strings.Builder
	total := fmt.Sprintf("%v", result)

	output.WriteString("\n\nCalculation:\n" + input1 + " " + operation + " " + input2 + " = " + total)

	return output.String()
}
