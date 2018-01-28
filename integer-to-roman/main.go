// https://leetcode.com/problems/integer-to-roman/description/
package main

import (
	"fmt"
	"strings"
)

var symbols = []map[int]string{
	{
		1: "I",
		4: "IV",
		5: "V",
		9: "IX",
	},
	{
		1: "X",
		4: "XL",
		5: "L",
		9: "XC",
	},
	{
		1: "C",
		4: "CD",
		5: "D",
		9: "CM",
	},
	{
		1: "M",
	},
}

func intToRoman(num int) string {
	s := ""
	mag := 0
	for num > 0 {
		d := num % 10
		switch {
		case d >= 1 && d <= 3:
			s = strings.Repeat(symbols[mag][1], d) + s
		case d == 4:
			s = symbols[mag][4] + s
		case d >= 5 && d <= 8:
			s = symbols[mag][5] + strings.Repeat(symbols[mag][1], d-5) + s
		case d == 9:
			s = symbols[mag][9] + s
		}
		num /= 10
		mag += 1
	}
	return s
}

func main() {
	q := []int{123, 2222, 1234, 1, 0, 22, 4, 44, 444, 3999}
	for _, v := range q {
		fmt.Printf("%d = %s\n", v, intToRoman(v))
	}
}
