package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	mx := max(len(a), len(b))
	var s []rune
	A := []rune(a)
	B := []rune(b)
	slices.Reverse(A)
	slices.Reverse(B)
	var carry int

	for i := 0; i < mx; i++ {
		sum := 0
		if i < len(A) && A[i] == '1' {
			sum++
		}
		if i < len(B) && B[i] == '1' {
			sum++
		}
		sum += carry
		if sum == 3 {
			s = append(s, '1')
			carry = 1
		} else if sum == 2 {
			s = append(s, '0')
			carry = 1
		} else if sum == 1 {
			s = append(s, '1')
			carry = 0
		} else {
			s = append(s, '0')
			carry = 0
		}
	}

	if carry > 0 {
		s = append(s, '1')
	}

	slices.Reverse(s)
	return string(s)
}
