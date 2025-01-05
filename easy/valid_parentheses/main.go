package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
}

type Stack struct {
	Values []byte
}

func (s *Stack) Push(ch byte) {
	s.Values = append(s.Values, ch)
}

func (s *Stack) Pop() byte {
	if len(s.Values) == 0 {
		return 0
	}

	ch := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]

	return ch
}

func isValid(s string) bool {
	stack := &Stack{}
	for _, ch := range s {
		if isOpenBracket(byte(ch)) {
			stack.Push(byte(ch))
			continue
		}

		el := stack.Pop()
		if el == 0 || !validCloseBracket(el, byte(ch)) {
			return false
		}
	}

	return len(stack.Values) == 0
}

func isOpenBracket(ch byte) bool {
	if ch == '(' || ch == '{' || ch == '[' {
		return true
	}

	return false
}

func validCloseBracket(el, ch byte) bool {
	switch ch {
	case ')':
		return el == '('
	case '}':
		return el == '{'
	case ']':
		return el == '['
	default:
		return false
	}
}
