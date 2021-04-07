package main

import "fmt"

func main() {
	var arr []int
	var d int
	var size int

	fmt.Printf("Enter the size of an array: ")
	fmt.Scanf("%d", &size)

	fmt.Println("Enter the number of an array: ")
	for i := 0; i < size; i++ {
		fmt.Scan(&d)

		arr = append(arr, d)
	}

	fmt.Println("Given array:", arr)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] == arr[j] {
				fmt.Println("Found duplicated number:", arr[i], "in positions:", i, "and", j)
			}
		}
	}

}
