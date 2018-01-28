// https://leetcode.com/problems/reverse-integer/description/
package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		x /= 10
		if result < math.MinInt32 || result > math.MaxInt32 {
			return 0
		}
	}
	return result
}

func main() {
	q := []int{
		123, -123, 120,
	}
	for _, v := range q {
		fmt.Printf("reverse %d = %d\n", v, reverse(v))
	}
}
