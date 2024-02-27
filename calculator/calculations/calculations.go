package calculations

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate(input1, input2, operation string) float64 {
	var value1 float64 = convertInputToValue(input1)
	var value2 float64 = convertInputToValue(input2)

	switch operation {
	case "+":
		return addValues(value1, value2)
	case "-":
		return subtractValues(value1, value2)
	case "*":
		return multiplyValues(value1, value2)
	case "/":
		return divideValues(value1, value2)
	default:
		panic("Invalid Operation")
	}
}

func convertInputToValue(input string) float64 {
	convertInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		message := fmt.Sprintf("%v must be a number", input)
		panic(message)
	}

	return convertInput
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
