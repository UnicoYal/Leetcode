package main

import "fmt"

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(16))
	fmt.Println(mySqrt(0))
}

func mySqrt(x int) int {
	left := 0
	right := x

	for left <= right {
		m := (left + right) / 2
		if m*m > x {
			right = m - 1
		} else if m*m < x {
			left = m + 1
		} else {
			return m
		}
	}

	return right
}
