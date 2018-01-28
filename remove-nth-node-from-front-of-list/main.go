package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromStart(head *ListNode, n int) *ListNode {
	if n == 1 {
		return head.Next
	}
	c := head
	for i := 0; i < n-2; i++ {
		c = c.Next
	}
	c.Next = c.Next.Next

	return head
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

func (n *ListNode) print() {
	if n == nil {
		fmt.Println()
		return
	}
	fmt.Printf("%d,", n.Val)
	n.Next.print()
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	tt := []struct {
		l []int
		n int
		e []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			1,
			[]int{2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{1, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
			[]int{1, 2, 3, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
			[]int{1, 2, 3, 4},
		},
	}
	for _, t := range tt {
		q := createList(t.l)
		r := removeNthFromStart(q, t.n)
		if !equal(fromList(r), t.e) {
			fmt.Println("expected:", t.e, "actual:", fromList(r))
			panic("failed case")
		}
	}
}
