package main

import (
	"fmt"
	"strings"
)

// use another function

func main() {

	PrintMathMethod()

	fmt.Println("\n\n\nGoodBye :)\n\n\n ")

}

func PrintMathMethod() {
	var num1, num2, result int
	var operator, IsMore string

	fmt.Println("write first number : ")
	fmt.Scanln(&num1)
	fmt.Println("write second number : ")
	fmt.Scanln(&num2)
	fmt.Println("write operator name (Plus,Mines,Multiplication,Division) : ")
	fmt.Scanln(&operator)

	result = CustomMath(num1, num2, operator)

	fmt.Println("\n for this operator :" + strings.ToUpper(operator) + ", result is : ")
	fmt.Println(result)

	fmt.Println("\ndo you want continue ? (yes or no) \n ")
	fmt.Scanln(&IsMore)

	if strings.ToLower(IsMore) == "yes" {
		PrintMathMethod()
	}
}

func CustomMath(x int, y int, operator string) int {

	switch strings.ToLower(operator) {
	case "plus":
		return x + y
	case "mines":
		return x - y
	case "multiplication":
		return x * y
	case "division":
		return x / y
	default:
		return 0
	}

}
