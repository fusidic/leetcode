package q2addtwonumbers

/**
*	You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
*
*	You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*
*	Example:
*
*	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
*	Output: 7 -> 0 -> 8
*	Explanation: 342 + 465 = 807.
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// type ListNode structures.ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	var flag int
	for true {
		p1.Val += flag
		p1.Val += p2.Val
		flag = p1.Val / 10
		p1.Val = p1.Val % 10
		if p1.Next == nil && p2.Next != nil {
			p1.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
		} else if p1.Next != nil && p2.Next == nil {
			p2.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
		} else if p1.Next == nil && p2.Next == nil {
			if flag == 0 {
				break
			}
			p1.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return l1
	// fmt.Printf(format string, a ...interface{})
}
