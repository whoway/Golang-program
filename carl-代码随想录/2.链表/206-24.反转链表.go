/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var sentry *ListNode
	for nil != head {
		temp := head.Next
		head.Next = sentry
		sentry = head
		head = temp
	}
	return sentry
}
