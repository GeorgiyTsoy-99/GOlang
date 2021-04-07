package main

import "fmt"

func input(x) []int {

	var d int
	n := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return input(x)
}

func main() {
	fmt.Println("Enter input:")
	x := input([]int{})
	fmt.Println("Input:", x)
}
