package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// link: https://leetcode.com/problems/linked-list-cycle
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	//
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
