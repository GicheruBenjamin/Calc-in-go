package main

import (
	"fmt"
)

func main() {
	fmt.Println("A calculator")
	fmt.Println("Please choose what you want to do:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Modulus")
	fmt.Println("6. Exit")

	var input int
	fmt.Scanln(&input)

	if input == 1 {
		fmt.Println("Please enter the first number:")
		var num1 int
		fmt.Scanln(&num1)
		fmt.Println("Please enter the second number:")
		var num2 int
		fmt.Scanln(&num2)
		fmt.Println("The result is:", addNumbers(num1, num2))
	} else if input == 2 {
		fmt.Println("Please enter the first number:")
		var num1 int
		fmt.Scanln(&num1)
		fmt.Println("Please enter the second number:")
		var num2 int
		fmt.Scanln(&num2)
		fmt.Println("The result is:", subtractNumbers(num1, num2))
	} else if input == 3 {
		fmt.Println("Please enter the first number:")
		var num1 int
		fmt.Scanln(&num1)
		fmt.Println("Please enter the second number:")
		var num2 int
		fmt.Scanln(&num2)
		fmt.Println("The result is:", multiplyNumbers(num1, num2))
	} else if input == 4 {
		fmt.Println("Please enter the first number:")
		var num1 int
		fmt.Scanln(&num1)
		fmt.Println("Please enter the second number:")
		var num2 int
		fmt.Scanln(&num2)
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		} else {
			fmt.Println("The result is:", divideNumbers(num1, num2))
		}
	} else if input == 5 {
		fmt.Println("Please enter the first number:")
		var num1 int
		fmt.Scanln(&num1)
		fmt.Println("Please enter the second number:")
		var num2 int
		fmt.Scanln(&num2)
		fmt.Println("The result is:", modulusNumbers(num1, num2))
	} else if input == 6 {
		fmt.Println("Exiting...")
		return
	} else {
		fmt.Println("Invalid input")
	}
}

func addNumbers(num1, num2 int) int {
	return num1 + num2
}

func subtractNumbers(num1, num2 int) int {
	return num1 - num2
}

func multiplyNumbers(num1, num2 int) int {
	return num1 * num2
}

func divideNumbers(num1, num2 int) int {
	return num1 / num2
}

func modulusNumbers(num1, num2 int) int {
	return num1 % num2
}
