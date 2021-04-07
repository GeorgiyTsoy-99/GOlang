package main

import (
	"fmt"
)

func main() {
	var originalInput string
	var status int = 0
	fmt.Println("Enter word to check whether it Palindrome or not")
	fmt.Scanf("%s", &originalInput)
	for i := 0; i < len(originalInput)/2; i++ {
		if originalInput[i] != originalInput[len(originalInput)-i-1] {
			status = 1
		}
	}

	if status == 1 {
		fmt.Println(originalInput, "is not a Palindrome")
	} else {
		fmt.Println(originalInput, "is a Palindrome")
	}
}
