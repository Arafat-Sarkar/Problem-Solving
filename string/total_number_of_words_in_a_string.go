package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your String")

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	// for i := 0; i < len(s); i++ {
	// 	ch:= string(s[i])
	// 	fmt.Println(ch)
	// }

	fmt.Println(len(s))

}
