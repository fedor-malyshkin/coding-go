package p84

/**
https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return helper(0, head.Val, head.Next)
}

func helper(repeatedIndex int, prevVal int, node *ListNode) *ListNode {
	if node == nil {
		if repeatedIndex > 0 {
			return nil
		} else {
			return &ListNode{prevVal, nil}
		}
	}
	if prevVal != node.Val {
		if repeatedIndex > 0 {
			return helper(0, node.Val, node.Next)
		} else {
			return &ListNode{prevVal, helper(0, node.Val, node.Next)}
		}
	} else {
		return helper(repeatedIndex+1, prevVal, node.Next)
	}
}
