package main

import (
	"SimpleCalculator/pkg/calculator"
	"fmt"
	"io"
	"os"
	"strings"
)

func NewCalc(reader io.Reader, calc calculator.Simple) {
	fmt.Println("Welcome to the Simple Calculator! 🧮")
	fmt.Println()
	fmt.Println("\tUsage: Any numbers you write will be added together")
	fmt.Println("\t\tEx: 1,2,3,4\t will return 10")
	fmt.Println("\t\tEx: 1,2,3,4,5,6,7,8,9,10\t will return 55")
	fmt.Println("\tYou can also use a custom separator")
	fmt.Println("\t\tEx: //;<newline>1;2;3;4;5;6;7;8;9;10\t will return 55")
	fmt.Println("\tStandard separators are: , and newline")
	fmt.Println()
	fmt.Print("scalc: ")

	// Loop to get input from the user exit when two newlines are entered
	var calcString strings.Builder

	// loop over input until two newlines are entered
	for {
		var input string
		_, err := fmt.Fscanln(reader, &input)
		if err != nil && err.Error() == "unexpected newline" {
			break
		} else if err != nil {
			panic("Input handling broke: " + err.Error())
		}

		calcString.WriteString(input + "\n")
	}

	// Calculate the result
	result, err := calc.Add(calcString.String())
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The result is", result)
	}
}

func main() {
	reader := os.Stdin
	logger := calculator.NewSimpleLogger()
	calc := calculator.NewSimple(logger)

	NewCalc(reader, calc)
}
