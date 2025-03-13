package main

import (
	"fmt"
)

func main() {

	var number1, number2 float64

	var operator string

	fmt.Print("give me the first value for calculation : ")
	fmt.Scanln(&number1)

	fmt.Print("give me the operator for calculation (+,-,*,/,%) : ")
	fmt.Scanln(&operator)

	fmt.Print("give me the second number for calculation : ")
	fmt.Scanln(&number2)

	switch operator {

	case "+":
		{
			fmt.Println("Addition")
			fmt.Printf("%.1f %s %.1f = %.1f", number1, operator, number2, number1+number2)
		}
	case "-":
		{
			fmt.Println("Subraction")
			fmt.Printf("%.1f %s %.1f = %.1f", number1, operator, number2, number1-number2)
		}
	case "*":
		{
			fmt.Println("Multiply")
			fmt.Printf("%.1f %s %.1f = %.1f", number1, operator, number2, number1*number2)
		}
	case "/":
		{
			fmt.Println("Division")
			if number2 == 0.0 {
				fmt.Println("you cant divide an number by zero bro")
			} else {
				fmt.Printf("%.1f %s %.1f = %.1f", number1, operator, number2, number1/number2)
			}
		}
	case "%":
		{
			fmt.Println("Reminder")
			if number2 == 0.0 {
				fmt.Println("you cant divide an number by zero bro")
			} else {
				fmt.Printf("%d %s %d = %d", int(number1), operator, int(number2), int(number1)%int(number2))
			}
		}
	default:
		{
			fmt.Println("Invalid operators or values check it back")
		}
	}

}
