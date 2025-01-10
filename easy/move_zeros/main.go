package main

import "fmt"

func main() {
	sl1 := []int{0,1,0,3,12}
	sl2 := []int{0}
	sl3 := []int{1,0,3,0,12,0}
	moveZeroes(sl1)
	moveZeroes(sl2)
	moveZeroes(sl3)

	fmt.Println(sl1)
	fmt.Println(sl2)
	fmt.Println(sl3)
}

func moveZeroes(nums []int) {
	var left int
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
