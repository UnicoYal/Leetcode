package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))

	fmt.Println("Solution 2: ")
	fmt.Println(lengthOfLastWord2("Hello World"))
	fmt.Println(lengthOfLastWord2("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord2("luffy is still joyboy"))
}

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	var wasLet bool
	var counter int

	for i >= 0 && (s[i] != ' ' || !wasLet) {
		if s[i] != ' ' {
			wasLet = true
			counter++
		}
		i--
	}

	return counter
}

func lengthOfLastWord2(s string) int {
	i := len(s) - 1
	var counter int
	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		counter++
		i--
	}

	return counter
}
