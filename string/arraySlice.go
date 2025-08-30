package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[2:4] //array thekhe slice
	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	s1 := s[1:2]
	fmt.Println(s1)
}
