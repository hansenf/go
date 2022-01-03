package main

import "fmt"

func IsOdd(num int) bool {
	if num <= 0 {
		return false
	} else if num%2 == 0 {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println(IsOdd(0))
}