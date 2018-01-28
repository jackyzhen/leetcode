// https://leetcode.com/problems/remove-linked-list-elements/description/
package main

import "fmt"

/**
 * Definition for singly-linked list.
 **/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	pseudoHead := &ListNode{
		0,
		head,
	}
	prev := pseudoHead
	curr := head

	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return pseudoHead.Next
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
		q []int
		d int
		e []int
	}{
		{
			[]int{1, 2, 2, 2, 1, 4, 5, 6, 2, 6},
			2,
			[]int{1, 1, 4, 5, 6, 6},
		},
		{
			[]int{1},
			1,
			[]int{},
		},
		{
			[]int{1, 1},
			1,
			[]int{},
		},
		{
			[]int{2},
			1,
			[]int{2},
		},
		{
			[]int{2, 1},
			1,
			[]int{2},
		},
		{
			[]int{2, 1},
			2,
			[]int{1},
		},
	}
	for _, t := range tt {
		q := createList(t.q)
		r := removeElements(q, t.d)
		if !equal(fromList(r), t.e) {
			fmt.Printf("removed %d. Expected %v actual %v\n", t.d, t.e, fromList(r))
			panic("failed case")
		}
	}
}
