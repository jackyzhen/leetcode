// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
package main

import "fmt"

func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i += 1
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	tt := []struct {
		nums     []int
		expected []int
		l        int
	}{
		{
			[]int{-1, 0, 0, 0, 0, 3, 3},
			[]int{-1, 0, 3},
			3,
		},
		{
			[]int{1, 1, 2},
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 1, 1, 2},
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 1, 2, 2},
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 1, 1, 1},
			[]int{1},
			1,
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3, 4},
			4,
		},
		{
			[]int{1, 1, 2, 3, 4},
			[]int{1, 2, 3, 4},
			4,
		},
		{
			[]int{1, 2, 3, 4, 4},
			[]int{1, 2, 3, 4},
			4,
		},
		{
			[]int{1, 2, 3, 3, 4, 4, 5},
			[]int{1, 2, 3, 4, 5},
			5,
		},
	}
	for _, t := range tt {
		actual := removeDuplicates(t.nums)
		fmt.Println(t.l, actual, t.expected, t.nums)
		for i := 0; i < t.l; i++ {
			if t.nums[i] != t.expected[i] {
				fmt.Printf("Expected len %d arr %v actual %v\n", t.l, t.expected, t.nums)
				panic("failed case")
			}
		}
	}
}
