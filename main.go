package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Public Vars
var title, endl string = "Simple Calculator", "------------------------"

func createCalc() {
	fmt.Println(endl, title, endl)

	calc := `
 ______________
|_____CAL______|
|___________=01|
| + - * / %  AC|
| 1 2 3 4 5 DEL|
| 6 7 8 9 0 ANS|
|        ON/OFF|
 --------------
	`
	fmt.Println(calc)
}

func clearScreen() {
	screen := exec.Command("clear")
	screen.Stdout = os.Stdout
	screen.Run()
}

func pickOperation(operation string) {
	clearScreen()
	operation = strings.ToLower(operation)

	if operation == "add" || operation == "addition" {
		addition()
	} else if operation == "sub" || operation == "subtraction" {
		subtraction()
	} else if operation == "mul" || operation == "multiplication" {
		multiplication()
	} else if operation == "div" || operation == "division" {
		division()
	} else if operation == "tab" || operation == "table" {
		printMultiplicationTable()
	} else {
		fmt.Println("Please, Enter a valid operation!")
	}
}

func main() {
	// cleas everything at start
	clearScreen()

	createCalc()

	// infinite loop started
	for {
		// pick operation
		fmt.Println("type: exit to quit the process..")
		fmt.Println("Please, select a operation: (add, sub, mul, div, tab)")
		var userInput string
		fmt.Scanln(&userInput)

		pickOperation(userInput)

		if userInput == "exit" || userInput == "EXIT" {
			break
		}
	}
}

// Operations
func addition() {
	fmt.Println("------ Addition In Progress ------")
	fmt.Println("How many values would you like to add?")
	var totalValues int
	fmt.Scanln(&totalValues)

	// final result
	var additionTotal int

	for i := 1; i < totalValues+1; i++ {
		fmt.Println("Enter value", i)
		var temp int
		fmt.Scanln(&temp)

		additionTotal += temp
	}

	clearScreen()
	fmt.Println("Total of Addition:", additionTotal)
}

func subtraction() {
	fmt.Println("------ Subtraction In Progress ------")
	var val1, val2 int

	fmt.Println("Enter Value 1:")
	fmt.Scanln(&val1)

	fmt.Println("Enter Value 2 (this will be subtracted from previous value):")
	fmt.Scanln(&val2)

	// final result
	subtractedTotal := val1 - val2
	clearScreen()
	fmt.Println("Total of Subtraction:", subtractedTotal)
}

func multiplication() {
	fmt.Println("------ Multiplication In Progress ------")
	var val1, val2 int

	fmt.Println("Enter Value 1:")
	fmt.Scanln(&val1)

	fmt.Println("Enter Value 2:")
	fmt.Scanln(&val2)

	// final result
	multiplicationTotal := val1 * val2
	clearScreen()
	fmt.Println("Total of Multiplication:", multiplicationTotal)
}

func division() {
	fmt.Println("------ Division In Progress ------")
	var dividend, divisor int

	fmt.Println("Enter Dividend:")
	fmt.Scanln(&dividend)

	fmt.Println("Enter Divisor")
	fmt.Scanln(&divisor)

	// final result
	quotient := dividend / divisor
	clearScreen()
	fmt.Println("Answer of Division:", quotient)
}

func printMultiplicationTable() {
	fmt.Println("------ Multiplication Tables In Progress ------")
	var tableNumber int

	fmt.Println("Which multiplication table would you like to print?")
	fmt.Scanln(&tableNumber)

	// final result
	clearScreen()
	fmt.Println(endl)
	for i := 1; i <= 10; i++ {
		fmt.Println(tableNumber, "x", i, "=", tableNumber*i)
	}
}
