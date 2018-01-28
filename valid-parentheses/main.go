// https://leetcode.com/problems/valid-parentheses/description/
package main

import "fmt"

type stack struct {
	arr []rune
}

func (s *stack) size() int {
	return len(s.arr)
}
func (s *stack) peek() rune {
	return s.arr[len(s.arr)-1]
}

func (s *stack) pop() rune {
	p := s.peek()
	s.arr = s.arr[:len(s.arr)-1]
	return p
}

func (s *stack) push(r rune) {
	s.arr = append(s.arr, r)
}

var op = map[rune]rune{
	'}': '{',
	')': '(',
	']': '[',
}

func isValid(s string) bool {
	st := stack{}
	for _, u := range s {
		switch u {
		case ')', '}', ']':
			if st.size() == 0 {
				return false
			}
			prev := st.peek()
			if prev != op[u] {
				return false
			}
			st.pop()
		case '(', '{', '[':
			st.push(u)
		default:
			return false
		}
	}
	return st.size() == 0
}

func main() {
	q := "({})(){}[(({{}}))]{}"
	fmt.Println(isValid(q))
}
