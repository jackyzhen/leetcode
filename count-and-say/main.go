// https://leetcode.com/problems/count-and-say/description

package main

import "fmt"

func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}

	return say(countAndSay(n - 1))
}

func say(num string) string {
	sayStr := ""
	count := 1
	for i := 0; i < len(num); i++ {
		if i == len(num)-1 {
			sayStr = fmt.Sprintf("%s%d%s", sayStr, count, string(num[i]))
			return sayStr
		}
		if num[i] != num[i+1] {
			sayStr = fmt.Sprintf("%s%d%s", sayStr, count, string(num[i]))
			count = 1
		} else {
			count += 1
		}
	}

	return sayStr
}

func main() {
	tt := []struct {
		test     int
		expected string
	}{
		{
			1, "1",
		},
		{
			2, "11",
		},
		{
			3, "21",
		},
		{
			4, "1211",
		},
		{
			5, "111221",
		},
		{
			6, "312211",
		},
		{
			7, "13112221",
		},
	}

	for _, t := range tt {
		actual := countAndSay(t.test)
		if t.expected != actual {
			panic(fmt.Sprintf("test: %v, expected %v, actual %v", t.test, t.expected, actual))
		}
	}
}
