/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var head  *ListNode
	var curr *ListNode
	for list1 != nil || list2 != nil {

		if list2 == nil || list1 != nil && list1.Val <= list2.Val {
			// take from list 1
			//fmt.Printf("Taking from list 1 val=%d\n", list1.Val)
			newNode := &ListNode{Val: list1.Val}
			if curr == nil {
				curr = newNode
				head = curr
			} else {
				curr.Next = newNode
				curr = curr.Next
			}
			list1 = list1.Next
		} else {
			// take from list 2
			//fmt.Printf("Taking from list 2 val=%d\n", list2.Val)
			newNode := &ListNode{Val: list2.Val}
			if curr == nil {
				curr = newNode
				head = curr
			} else {
				curr.Next = newNode
				curr = curr.Next
			}
			list2 = list2.Next
		}
	}
	return head
}
