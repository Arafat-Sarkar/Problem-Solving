package main

import "fmt"

func PrintNum(n int) {
	if(n==0){
		return
	}
	fmt.Println(n)
	PrintNum(n-1)
}

func main() {
	PrintNum(5)

}