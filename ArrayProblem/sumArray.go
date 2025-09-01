package main

import "fmt"

func main() {

	var n int
	sum := 0
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		sum = sum + arr[i]
	}

	fmt.Println(sum)

}
