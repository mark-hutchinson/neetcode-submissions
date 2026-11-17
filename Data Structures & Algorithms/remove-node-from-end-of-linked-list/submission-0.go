/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}

	target := count - n
    // fmt.Printf("Target to delete %d\n", target)
	count = 0
	var prev *ListNode
	curr = head
	for count <= target{
		if count == target {
			if prev == nil {
				// pop off first element - nothing previous to skip
				head = curr.Next
			} else {
				// Delete curr, by skipping over it
				prev.Next = curr.Next
			}
			break
		}
		count++
		prev = curr
		curr = curr.Next
	}

	return head
}
