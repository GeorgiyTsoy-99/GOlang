package main

import (
	"fmt"
)

func main() {
	var number int
	var evenSum int = 0
	var oddSum int = 0

	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &number)

	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}
	fmt.Println("Sum of even numbers:", evenSum)
	fmt.Println("Sum of odd numbers:", oddSum)
}
