// https://leetcode.com/problems/longest-palindromic-substring/description/

package main

import "fmt"

func longestPalindrome(s string) string {
	chars := []rune(s)
	n := len(chars)
	if n <= 1 {
		return s
	}
	size := n
	for size > 1 {
		for i := 0; i+size <= n; i++ {
			if isPalindrome(chars[i : size+i]) {
				return string(chars[i : size+i])
			}
		}
		size -= 1
	}
	return string(chars[0])
}

func isPalindrome(chars []rune) bool {
	left := 0
	right := len(chars) - 1
	for left < right {
		if chars[left] != chars[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func main() {
	q := []string{
		"babad",
		"cbbd",
		"dasdasghghghghghi3uajsf",
	}
	for _, s := range q {
		fmt.Printf("%s = %s\n", s, longestPalindrome(s))
	}
}
