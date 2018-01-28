// https://leetcode.com/problems/two-sum/description/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		comp := target - nums[i]
		if ind, ok := m[comp]; ok {
			return []int{i, ind}
		}
		m[nums[i]] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
