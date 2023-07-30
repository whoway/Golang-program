/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	// slice来模拟栈
	var res []int
	for nil != head {
		res = append(res, head.Val)
		head = head.Next
	}
	mylen := len(res)
	Left := 0
	Right := mylen - 1
	for Left < Right {
		res[Left] ^= res[Right]
		res[Right] ^= res[Left]
		res[Left] ^= res[Right]
		Left++
		Right--
	}
	return res
}
