package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Sonni kiriting: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("togri son kiritng !")
		return
	}

	fmt.Println("2 dan", n, "gacha bo'lgan tub sonlar:")

	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func IsPrime(num int) bool {
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
