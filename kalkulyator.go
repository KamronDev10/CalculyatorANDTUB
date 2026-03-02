package main

import (
	"errors"
	"fmt"
)

func main() {
	var operator string

	fmt.Println("To'xtatish uchun amal o'rniga 'q' ni kiriting.")

	for {
		var num1, num2 float64

		fmt.Print("Birinchi sonni kiriting: ")
		value, err := fmt.Scan(&num1)
		if err != nil {
			fmt.Println("Xato: Son kiriting!")
			return
		}
		fmt.Println(value)
		fmt.Print("Amalni kiriting (+, -, *, /) yoki 'q': ")
		fmt.Scan(&operator)

		if operator == "q" {
			break
		}

		fmt.Print("Ikkinchi sonni kiriting: ")
		_, err = fmt.Scan(&num2)
		if err != nil {
			fmt.Println("Xato: Son kiriting!")
			continue
		}

		var result float64
		var calculationError error

		switch operator {
		case "+":
			result = Add(num1, num2)
		case "-":
			result = Sub(num1, num2)
		case "*":
			result = Mul(num1, num2)
		case "/":
			result, calculationError = Div(num1, num2)
		default:
			fmt.Println("Xato: Noto'g'ri amal!")
			continue
		}

		if calculationError != nil {
			fmt.Println("Xato:", calculationError)
			continue
		}

		fmt.Println("Natija:", num1, operator, num2, "=", result)
	}

	fmt.Println("Dastur yopildi.")
}

func Add(a, b float64) float64 { return a + b }
func Sub(a, b float64) float64 { return a - b }
func Mul(a, b float64) float64 { return a * b }
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Nolga bo'lish mumkin emas!")
	}
	return a / b, nil
}
