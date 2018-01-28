// https://leetcode.com/problems/3sum-closest/description/
package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	if n < 3 {
		return 0
	}
	ans := nums[0] + nums[1] + nums[n-1]
	for i := 0; i < n; i++ {
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			ansDiff := abs(ans - target)
			sumDiff := abs(sum - target)
			if ansDiff > sumDiff {
				ans = sum
				fmt.Println(nums[i], nums[j], nums[k])
			}
			if sum < target {
				j += 1
			} else if sum > target {
				k -= 1
			} else {
				j, k = j+1, k-1
			}
		}
	}
	return ans

}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func main() {

	q := []int{-1, 2, 1, -4, 3, 1, 12}
	t := 50
	fmt.Println(threeSumClosest(q, t))
}
