package main

import "fmt"

func main() {
	sl1 := []int{1, 3, 5, 6}
	sl2 := []int{1, 3, 5, 6}
	sl3 := []int{1, 3, 5, 6}

	t1 := 5
	t2 := 2
	t3 := 7

	fmt.Println(searchInsert(sl1, t1))
	fmt.Println(searchInsert(sl2, t2))
	fmt.Println(searchInsert(sl3, t3))
}

func searchInsert(nums []int, target int) int {
	var left int
	right := len(nums) - 1

	for left < right {
		m := (right + left) / 2
		if nums[m] > target {
			right = m - 1
		} else if nums[m] < target {
			left = m + 1
		} else {
			return m
		}
	}

	if nums[right] < target {
		return right + 1
	}

	if right == 0 {
		return 0
	}
	return right - 1
}
