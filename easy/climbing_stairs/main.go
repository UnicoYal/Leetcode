package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}

func climbStairs(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
			a, b = b, a+b
	}

	return b
}
