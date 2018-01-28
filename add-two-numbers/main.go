//https://leetcode.com/problems/add-two-numbers/

package main

import "fmt"

/**
 * Definition for singly-linked list.
 **/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans *ListNode
	var curr *ListNode
	var remainder int
	for l1 != nil || l2 != nil || remainder != 0 {
		var val int
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		val += remainder
		remainder = val / 10

		nn := ListNode{
			Val: val % 10,
		}
		if ans == nil {
			ans = &nn
		}
		if curr == nil {
			curr = &nn
		} else {
			curr.Next = &nn
			curr = &nn
		}
	}
	return ans
}

func createList(vals []int) *ListNode {
	var h *ListNode
	c := h
	for _, v := range vals {
		if c == nil {
			h = &ListNode{Val: v}
			c = h
		} else {
			c.Next = &ListNode{Val: v}
			c = c.Next
		}
	}
	return h
}

func fromList(h *ListNode) []int {
	arr := []int{}
	for h != nil {
		arr = append(arr, h.Val)
		h = h.Next
	}
	return arr
}

func main() {
	a := createList([]int{2, 4, 3})
	b := createList([]int{5, 6, 4})
	ans := addTwoNumbers(a, b)
	fmt.Println(fromList(ans))
}
