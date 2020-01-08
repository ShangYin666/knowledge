package main

// link: https://leetcode.com/problems/reverse-linked-list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
1->2->3->4->5->nil
5->4->3->2->1->nil
循环每次将前后节点进行交换
第一种比第二种方法在内存方面更夹节省，每次循环少申请一个零时变量next
*/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
