package main

import "fmt"

func main() {
	sl1 := []int{1, 2, 3}
	sl2 := []int{4, 3, 2, 1}
	sl3 := []int{9}
	fmt.Println(plusOne(sl1))
	fmt.Println(plusOne(sl2))
	fmt.Println(plusOne(sl3))

	fmt.Println("Second solution: ")
	sl1 = []int{1, 2, 3}
	sl2 = []int{4, 3, 2, 1}
	sl3 = []int{9}

	fmt.Println(plusOne2(sl1))
	fmt.Println(plusOne2(sl2))
	fmt.Println(plusOne2(sl3))
}

func plusOne(digits []int) []int {
	i := len(digits) - 1
	var ost int

	for i >= 0 {
		ost = (digits[i] + 1) / 10
		digits[i] = (digits[i] + 1) % 10
		if ost == 0 {
			break
		}
		i--
	}

	if ost == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	digits = append([]int{1}, digits...)
	return digits
}
