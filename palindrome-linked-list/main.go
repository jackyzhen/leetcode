// https://leetcode.com/problems/palindrome-linked-list/description/
package main

import "fmt"

/**
 * Definition for singly-linked list.
 **/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var (
		fast = head
		slow = head
		rev  *ListNode
	)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		temp := rev
		rev, slow = slow, slow.Next
		rev.Next = temp
	}
	if fast != nil {
		slow = slow.Next
	}
	for rev != nil {
		if rev.Val != slow.Val {
			return false
		}
		rev, slow = rev.Next, slow.Next
	}
	return true
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
		e bool
	}{
		{
			[]int{},
			true,
		},
		{
			[]int{3, 2, 1},
			false,
		},
		{
			[]int{1},
			true,
		},
		{
			[]int{1, 1},
			true,
		},
		{
			[]int{2, 1},
			false,
		},
		{
			[]int{2, 1, 2},
			true,
		},
	}
	for _, t := range tt {
		q := createList(t.q)
		r := isPalindrome(q)
		if r != t.e {
			fmt.Printf("%v actual %v, got %v\n", t.q, t.e, r)
			panic("failed case")
		}
	}
}
