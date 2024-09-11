// This is just a simple calculator for the cli it was a project to get started in go
package main

import (
	"fmt"
	"os"
)

func main() {
	var input int
	var number1 float64
	var number2 float64
	var answer string

	fmt.Println()
	fmt.Println("What math you wanna do (0 = exit)")
	fmt.Println("1. Adding")
	fmt.Println("2. Subtraktion")
	fmt.Println("3. Multipling")
	fmt.Println("4. Dividing")
	fmt.Scan(&input)
	if input == 0 {
		os.Exit(0)
	}
	fmt.Println()
	fmt.Println("Your two numbers plz:")
	fmt.Scanln(&number1, &number2)

	if input == 1 {
		answer = adding(number1, number2)

	} else if input == 2 {
		answer = subtraktion(number1, number2)

	} else if input == 3 {
		answer = multiplying(number1, number2)

	} else if input == 4 {
		answer = dividing(number1, number2)

	} else {
		fmt.Println("Thats not a valid command")
		main()
	}

	fmt.Println(answer)
	main()

}

func adding(number1 float64, number2 float64) string {
	solution := number1 + number2
	var answer string = fmt.Sprintf("%v + %v = %v", number1, number2, solution)
	return answer
}

func subtraktion(number1 float64, number2 float64) string {
	solution := number1 - number2
	var answer string = fmt.Sprintf("%v - %v = %v", number1, number2, solution)
	return answer
}

func multiplying(number1 float64, number2 float64) string {
	solution := number1 * number2
	var answer string = fmt.Sprintf("%v * %v = %v", number1, number2, solution)
	return answer
}

func dividing(number1 float64, number2 float64) string {
	solution := number1 / number2
	var answer string = fmt.Sprintf("%v / %v = %v", number1, number2, solution)
	return answer
}
