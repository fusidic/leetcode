package linklist

import (
	"container/list"

	struc "github.com/halfrost/LeetCode-Go/structures"
)

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	prev := hair
	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}

		next := tail.Next
		tail.Next = nil
		prev.Next = reverseList(head)
		// return reverseList(head)
		prev = head
		tail = prev
		// return prev
		head = next
		// return head
		prev.Next = head
	}

	return hair.Next
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func solution2(head *struc.ListNode, k int) *struc.ListNode {
	hair := &struc.ListNode{Next: head}
	stack := list.New()
	prev := hair

	for head != nil {
		for i := 0; i < k; i++ {
			stack.PushBack(head)
			head = head.Next
			if head == nil {
				return hair.Next
			}
		}

		for stack.Len() != 0 {
			prev.Next = stack.Back().Value.(*struc.ListNode)
			prev = prev.Next
			prev.Next = nil
			stack.Remove(stack.Back())
		}
		prev.Next = head
	}
	return hair.Next
}
