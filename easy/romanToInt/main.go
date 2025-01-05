package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	res := 0
	vals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && vals[s[i]] < vals[s[i+1]] {
			res -= vals[s[i]]
		} else {
			res += vals[s[i]]
		}
	}

	return res
}
