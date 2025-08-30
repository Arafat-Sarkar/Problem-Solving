package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your Full Name")

	s, _ := reader.ReadString('\n')
	// fmt.Println(strings.ToUpper(s))
	// fmt.Println(strings.Contains(s,"rifat"))
	if strings.Contains(s, "rifat") == true {
		fmt.Println("found")
	} else {
		fmt.Println("Not found")
	}
	
}
