// https://leetcode.com/problems/roman-to-integer/description/
package main

import "fmt"

var numerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	v := 0
	prev := 1<<31 - 1
	for _, c := range s {
		curr := numerals[string(c)]
		if curr > prev {
			v += curr
			v -= prev * 2
		} else {
			v += curr
		}
		prev = curr
	}
	return v
}
func main() {
	q := []string{
		"CXXIII",
		"MMCCXXII",
		"MCCXXXIV",
		"I",
		"XXII",
		"IV",
		"XLIV",
		"CDXLIV",
		"MMMCMXCIX",
	}
	for _, s := range q {
		fmt.Printf("%s = %d\n", s, romanToInt(s))
	}
}
