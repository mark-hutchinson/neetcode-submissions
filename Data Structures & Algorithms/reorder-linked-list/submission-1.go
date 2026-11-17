/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	// First, find the middle of the list
	// Split the list into two lists
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	first := head
	second := slow.Next
	slow.Next = nil

	printList(first)
	//printList(second)

	// Next, reverse the second list
	var prev *ListNode
	curr := second
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev // reverse direction
		prev = curr
		curr = tmp
	}
	second = prev
	//printList(second)

	// Finally, merge the two lists
    // 2 4
	// 8 6
	for first != nil && second != nil {
		firstNext := first.Next
		secondNext := second.Next

		second.Next = firstNext
		first.Next = second


		first = firstNext
		second = secondNext

	}
}


func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Print("\n")
}