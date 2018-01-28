// https://leetcode.com/problems/container-with-most-water/description/
package main

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0
	for left != right {
		a := min(height[left], height[right]) * (right - left)
		if max < a {
			max = a
		}
		if height[left] > height[right] {
			right -= 1
		} else {
			left += 1
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
