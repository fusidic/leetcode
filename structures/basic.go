// Package structures provides structures that can be used to create various data structures.
package structures

// ListNode is a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeStack is the stack to store Treenode
type TreeStack struct {
	root TreeNode
	len  int
}
