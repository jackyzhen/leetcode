// https://leetcode.com/problems/reverse-linked-list/

package main

import "fmt"

/**
 * Definition for singly-linked list.
 **/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	h, _ := recurse(head)
	return h
}

func recurse(curr *ListNode) (*ListNode, *ListNode) {
	if curr == nil {
		return nil, nil
	}
	if curr.Next == nil {
		return curr, curr
	}
	head, currHead := recurse(curr.Next)
	curr.Next = nil
	currHead.Next = curr
	return head, curr
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
		e []int
	}{
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 1},
			[]int{1, 1},
		},
		{
			[]int{2, 1},
			[]int{1, 2},
		},
	}
	for _, t := range tt {
		q := createList(t.q)
		r := reverseList(q)
		if !equal(fromList(r), t.e) {
			fmt.Printf("Expected %v actual %v\n", t.e, fromList(r))
			panic("failed case")
		}
	}
}
