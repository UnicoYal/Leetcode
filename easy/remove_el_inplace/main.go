package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))
}

func removeElement(nums []int, val int) int {
	left := 0
	right := len(nums) - 1

	for left != right {
			if nums[left] == val {
					if nums[right] == val {
							right--
					} else {
							nums[left] = nums[right]
							left++
							right--
					}
					continue
			}
			left++
	}

	return left+1
}
