/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	prev := head
	curr := head.Next
	head.Next = nil
	for curr != nil {
		next := curr.Next // capture next node before we change it
		curr.Next = prev // reverse direction
		// advance the pointers
		prev = curr
		curr = next

	}
	return prev
}
