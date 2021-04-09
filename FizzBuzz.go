package main

import "fmt"

func main() {
	fmt.Print("Enter number: ")
	var num int
	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
	}
}
