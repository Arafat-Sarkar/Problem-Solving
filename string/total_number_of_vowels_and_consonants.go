package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	countVowel := 0
	countConstant := 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your String")

	s, _ := reader.ReadString('\n')

	for i := 0; i < len(s); i++ {
		ch := string(s[i])
		if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
			countVowel++
		} else if ch >= "a" && ch <= "z" {
			countConstant++
		}
	}

	fmt.Println("Vowel :", countVowel)
	fmt.Println("Constant :", countConstant)
}
