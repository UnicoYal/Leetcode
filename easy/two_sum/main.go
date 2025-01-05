package main

import "fmt"

func main() {
	firstExample := []int{2, 7, 11, 15}
	firstTarget := 9
	secondExample := []int{3, 2, 4}
	secondTarget := 6
	thirdExample := []int{3, 3}
	thirdTarget := 6

	fmt.Println(twoSum(firstExample, firstTarget))
	fmt.Println(twoSum(secondExample, secondTarget))
	fmt.Println(twoSum(thirdExample, thirdTarget))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 10)
	res := make([]int, 0, 2)
	for i, el := range nums {
		left := target - el
		if ind, ok := m[left]; ok {
			res = append(res, i, ind)
			break
		}
		m[el] = i
	}

	return res
}
