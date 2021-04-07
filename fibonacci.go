package main

import "fmt"

func main() {
	firstNum := 0
	secondNum := 1

	fmt.Println(firstNum)
	fmt.Println(secondNum)

	for i := 0; i < 10; i++ {
		sum := firstNum + secondNum
		fmt.Println(sum)
		firstNum = secondNum
		secondNum = sum
	}
}
