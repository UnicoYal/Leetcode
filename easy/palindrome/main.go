package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	newNum := 0
	curr := x
	for curr > 0 {
		left := curr % 10
		newNum = newNum*10 + left
		curr = curr / 10
	}

	return newNum == x
}
