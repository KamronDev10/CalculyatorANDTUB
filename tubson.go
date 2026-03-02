package main

import "fmt"

func main() {
	var n int

	fmt.Print("Son kiriting: ")

	count, err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("Xatolik yuz berdi:", err)
		return
	}

	fmt.Print(count)
	fmt.Print(", Kiritilgan son: ")
	fmt.Println(n)

	fmt.Print("2 dan ")
	fmt.Print(n)
	fmt.Println(" gacha bo'lgan tub sonlar:")

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
