package main

import "fmt"

func main() {

	s := make([]int, 3,5)
	fmt.Print(s )
	fmt.Print(len(s))
	fmt.Print(cap(s))
	s[2]=6
	fmt.Print(s )
}