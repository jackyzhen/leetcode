// https://leetcode.com/problems/valid-palindrome/description/
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func isPalindrome(s string) bool {
	if utf8.RuneCountInString(s) <= 1 {
		return true
	}
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; {
		if !isValid(chars[i]) {
			i += 1
			continue
		}
		if !isValid(chars[j]) {
			j -= 1
			continue
		}
		if !isSame(chars[i], chars[j]) {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}
func isValid(c rune) bool {
	return (c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z') ||
		(c >= '0' && c <= '9')
}

func isSame(a, b rune) bool {
	return strings.ToLower(string(a)) == strings.ToLower(string(b))
}
func main() {
	q := []string{
		"abcba",
		"abcdbb",
		"a",
		"",
		"aa",
		"aaa",
		"aab",
		"aba",
		"a.b.a",
		"90 1a    ()*@#$%^ b   . . . .c ... ()$&!*$& b a 10 $**($9",
	}
	for _, s := range q {
		fmt.Printf("%s %v\n", s, isPalindrome(s))
	}
}
