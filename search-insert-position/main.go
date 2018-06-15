// https://leetcode.com/problems/search-insert-position

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i, n := range nums {
		if n >= target {
			return i
		}
	}
	return len(nums)
}

func main() {
	tt := []struct {
		nums   []int
		target int
		e      int
	}{
		{
			[]int{1, 3, 5, 6}, 5, 2,
		},
		{
			[]int{1, 3, 5, 6}, 2, 1,
		},
		{
			[]int{1, 3, 5, 6}, 7, 4,
		},
		{
			[]int{1, 3, 5, 6}, 0, 0,
		},
	}

	for _, t := range tt {
		actual := searchInsert(t.nums, t.target)
		if t.e != actual {
			panic(fmt.Sprintf("nums: %v, target: %v, expected %v, actual %v", t.nums, t.target, t.e, actual))
		}
	}
}
