// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	max := 0
	chars := []rune(s)
	n := len(chars)
	if n < 2 {
		return n
	}

	m := map[rune]struct{}{}

	for i, j := 0, 0; i < n && j < n; {
		if _, ok := m[chars[j]]; ok {
			delete(m, chars[i])

			i += 1
		} else {
			m[chars[j]] = struct{}{}
			j += 1
		}
		if max < j-i {
			max = j - i
		}
	}
	return max
}

func main() {
	q := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}
	for _, s := range q {
		fmt.Println(lengthOfLongestSubstring(s))
	}

}
