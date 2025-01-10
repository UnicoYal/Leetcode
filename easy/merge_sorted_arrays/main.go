package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 && n != 0 {
		copy(nums1, nums2)
		return
	}

	A := m - 1
	B := n - 1
	C := m + n - 1

	for B >= 0 && A >= 0 {
		if nums1[A] > nums2[B] {
			nums1[C] = nums1[A]
			A--
		} else {
			nums1[C] = nums2[B]
			B--
		}
		C--
	}
}
