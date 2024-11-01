package main

import (
	"fmt"
)

func main(){
	fmt.println("A calculator")
	fmt.println("Please choose what U wanna do %n %n 1. Add %n 2. Subtract %n 3. Multiply %n 4. Divide")
	var input int
	fmt.Scanln(&input)
	if input == 1 {
		fmt.println("Please enter the first number")
		var num1 int
		fmt.Scanln(&num1)
		fmt.println("Please enter the second number")
		var num2 int
		fmt.Scanln(&num2)
		fmt.println("The result is: ", addNumbers(num1, num2))
	} else if input == 2 {
		fmt.println("Please enter the first number")
		var num1 int
		fmt.Scanln(&num1)
		fmt.println("Please enter the second number")
		var num2 int
		fmt.Scanln(&num2)
		fmt.println("The result is: ", subtractNumbers(num1, num2))
	} else if input == 3 {
		fmt.println("Please enter the first number")
		var num1 int
		fmt.Scanln(&num1)
		fmt.println("Please enter the second number")
		var num2 int
		fmt.Scanln(&num2)
		fmt.println("The result is: ", multiplyNumbers(num1, num2))
	}	
	else if input == 4 {
		fmt.println("Please enter the first number")
		var num1 int
		fmt.Scanln(&num1)
		fmt.println("Please enter the second number")
		var num2 int
		fmt.Scanln(&num2)
		fmt.println("The result is: ", divideNumbers(num1, num2))
	} else {
		fmt.println("Invalid input")
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