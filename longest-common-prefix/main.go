// https://leetcode.com/problems/longest-common-prefix/description/

package main

import "fmt"
import "unicode/utf8"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := []rune{}
	minN := minStr(strs)
	runes := make([][]rune, len(strs))
	for i := 0; i < len(runes); i++ {
		runes[i] = []rune(strs[i])
	}
	for i := 0; i < minN; i++ {
		char := runes[0][i]
		for j := 0; j < len(runes); j++ {
			if runes[j][i] != char {
				return string(prefix)
			}
		}
		prefix = append(prefix, char)
	}
	return string(prefix)
}

func minStr(strs []string) int {
	min := 1<<31 - 1
	for _, s := range strs {
		n := utf8.RuneCountInString(s)
		if n == 0 {
			return 0
		}
		if min > n {
			min = n
		}
	}
	return min
}
func main() {
	q := []string{
		"abcdefgaaaaaa",
		"abcdefgaaaaaaaaa",
		"abcdefg",
		"abcdehghgasd",
	}
	fmt.Println(longestCommonPrefix(q))
}
