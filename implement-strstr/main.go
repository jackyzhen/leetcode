// https://leetcode.com/problems/implement-strstr
package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	haystackRunes := []rune(haystack)
	needleRunes := []rune(needle)
	for i := 0; i < len(haystackRunes)-len(needleRunes)+1; i++ {
	innerLoop:
		for j := 0; j < len(needleRunes); j++ {
			if haystackRunes[i+j] != needleRunes[j] {
				break innerLoop
			}
			if j == len(needleRunes)-1 {
				return i
			}
		}
	}
	return -1
}

func main() {
	tt := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"", "", 0},
		{"aaa", "aaaa", -1},
		{"abc", "", 0},
		{"", "abc", -1},
		{"christmas", "ris", 2},
		{"christmaschristmas", "ris", 2},
		{"christmaschristmas", "christmas", 0},
		{"christmachristmas", "christmas", 8},
		{"xxxxxxxxxs", "s", 9},
		{"sxxxxxxxxx", "s", 0},
		{"mississippi", "issip", 4},
		{"mississippi", "issipi", -1},
	}

	for _, t := range tt {
		actual := strStr(t.haystack, t.needle)
		if actual != t.expected {
			fmt.Printf("Haystack %#v Needle %#v Expected %#v Actual %#v\n", t.haystack, t.needle, t.expected, actual)
		}
	}

}
