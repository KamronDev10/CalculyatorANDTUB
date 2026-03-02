package main

import (
	"fmt"
)

func main() {
	var num1, num2 int

	fmt.Print("Birinchi sonni kiriting: ")
	fmt.Scan(&num1)
	fmt.Print("Ikkinchi sonni kiriting: ")
	fmt.Scan(&num2)
	if num1 > num2 {
		num1, num2 = num2, num1
	}

	fmt.Println(num1, "va", num2, "orasidagi tub sonlar:")
	count := 0
	for i := num1; i <= num2; i++ {
		if isprime(i) {
			fmt.Print(i, " ")
			count++
		}
	}

	fmt.Println("\nJami tub sonlar soni:", count)
}

func isprime(num int) bool {
	if num < 2 {
		return false
	}
	for j := 2; j*j <= num; j++ {
		if num%j == 0 {
			return false
		}
	}
	return true
}
