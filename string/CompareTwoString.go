package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Your 1st String")

	s1, _ := reader.ReadString('\n')
	fmt.Println("Enter Your 2nd String")
	s2, _ := reader.ReadString('\n')

	if len(s1) == len(s2) {
		fmt.Println("Both are equal")
	} else if len(s1) > len(s2) {
		fmt.Println("Grater Then 1st String")
	} else {
		fmt.Println("Grater Then 2nd String")
	}

}
