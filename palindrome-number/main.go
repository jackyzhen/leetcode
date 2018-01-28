// https://leetcode.com/problems/palindrome-number/description/
package main

import "fmt"

func isPalindrome(x int) bool {
	if (x < 0) || (x%10 == 0 && x != 0) {
		return false
	}

	var rx int
	for x > rx {
		rx = rx*10 + x%10

		x /= 10
	}
	return x == rx || x == rx/10
}

func main() {
	q := []int{
		0, 1, 11, 121, 111, 1232, 2147483647, 2147447412, 33343, 33433,
	}
	for _, v := range q {
		fmt.Printf("isPalindrome('%d') = %v\n", v, isPalindrome(v))
	}
}
