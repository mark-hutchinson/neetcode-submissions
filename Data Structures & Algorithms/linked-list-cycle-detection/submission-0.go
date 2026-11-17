/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		// If theres a cycle, eventually slow and fast will meet
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
    return false
}
