package main

import (
	"fmt"
)

// Recursive function to find sum
func sumArray(arr []int, n int) int {
	// Base case
	if n == 0 {
		return 0
	}
	// Recursive case
	return arr[n-1] + sumArray(arr, n-1)
}

func main() {
	var n int
	fmt.Print("Enter size of array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Enter elements:")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	result := sumArray(arr, n)
	fmt.Println("Sum of array elements:", result)
}
