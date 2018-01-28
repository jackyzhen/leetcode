// https://leetcode.com/problems/string-to-integer-atoi/description/
package main

import (
	"fmt"
	"math"
	"strings"
)

const MAXINT = 1<<31 - 1
const MININT = ^MAXINT

var digits = map[string]uint{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func myAtoi(str string) int {
	result := uint(0)
	chars := []rune(strings.Trim(str, " "))
	neg := false
	if len(chars) == 0 {
		return 0
	}
	if string(chars[0]) == "-" {
		neg = true
		chars = chars[1:]
	} else if string(chars[0]) == "+" {
		chars = chars[1:]
	}
	for i := 0; i < len(chars); i++ {
		_, ok := digits[string(chars[i])]
		if !ok {
			chars = chars[:i]
		}
	}
	for i, c := range chars {
		d, _ := digits[string(c)]
		result += d * (uint(math.Pow10(len(chars) - i - 1)))
		if result > 1<<31-1 {
			if neg {
				return MININT
			} else {
				return MAXINT
			}
		}
	}

	if neg {
		return -int(result)
	} else {
		return int(result)
	}
}

func main() {
	q := []string{
		"123",
		"34.5",
		"wad12fa",
		"asdgsag",
		"2145fasf",
		"-9999999999999999999999",
		"9999999999999999999999",
	}
	for _, s := range q {
		fmt.Printf("atoi('%s') = %d\n", s, myAtoi(s))

	}

}
