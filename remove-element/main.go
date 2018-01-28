// https://leetcode.com/problems/remove-element/description/

package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i += 1
		}
	}
	return i
}

func main() {
	tt := []struct {
		nums     []int
		val      int
		expected []int
		l        int
	}{
		{
			[]int{-1, 0, 0, 0, 0, 3, 3},
			0,
			[]int{-1, 3, 3},
			3,
		},
		{
			[]int{1, 1, 2},
			1,
			[]int{2},
			1,
		},
		{
			[]int{1, 1, 1, 2},
			2,
			[]int{1, 1, 1},
			3,
		},
	}
	for _, t := range tt {
		actual := removeElement(t.nums, t.val)
		fmt.Println(t.l, actual, t.expected, t.nums)
		for i := 0; i < t.l; i++ {
			if t.nums[i] != t.expected[i] {
				fmt.Printf("Expected len %d arr %v actual %v\n", t.l, t.expected, t.nums)
				panic("failed case")
			}
		}
	}
}
