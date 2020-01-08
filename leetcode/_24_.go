package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// link : https://leetcode.com/problems/swap-nodes-in-pairs
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil {
		fast.Val, slow.Val = slow.Val, fast.Val
		slow = fast.Next
		if slow == nil {
			break
		}
		fast = fast.Next.Next
	}
	return head
}
