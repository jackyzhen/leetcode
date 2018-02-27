// https://leetcode.com/problems/generate-parentheses/

package main

import (
	"fmt"
	"reflect"
)

func genParams(n int) []string {
	w := []string{}
	recurse(n-1, n, 1, "(", &w)
	return w
}

func recurse(left, right, bal int, curr string, words *[]string) {
	if right == 0 {
		*words = append(*words, curr)
		return
	}
	if left == 0 {
		recurse(left, right-1, bal-1, curr+")", words)
	} else if bal > 0 {
		recurse(left-1, right, bal+1, curr+"(", words)
		recurse(left, right-1, bal-1, curr+")", words)
	} else {
		recurse(left-1, right, bal+1, curr+"(", words)
	}
}

func main() {
	tt := []struct {
		n int
		e []string
	}{
		{
			3,
			[]string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
		{
			2,
			[]string{
				"(())",
				"()()",
			},
		},
		{
			1,
			[]string{
				"()",
			},
		},
	}

	for _, t := range tt {
		actual := genParams(t.n)
		if !reflect.DeepEqual(actual, t.e) {
			panic(fmt.Sprintf("expected %#v, got %#v", t.e, actual))
		}
	}
}
