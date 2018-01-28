// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
package main

import (
	"fmt"
)

var pad = map[string]string{
	"1": "",
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var ans = []string{}

func letterCombinations(digits string) []string {
	pads := []string{}
	for _, c := range digits {
		pads = append(pads, pad[string(c)])
	}
	comb("", pads)
	return ans
}

func comb(l string, strings []string) {
	if len(strings) == 0 {
		if l != "" {
			ans = append(ans, l)
		}
		return
	}
	for i := 0; i < len(strings[0]); i++ {
		comb(l+string(strings[0][i]), strings[1:])
	}
}

func main() {
	a := "78"
	fmt.Println(letterCombinations(a))
}
