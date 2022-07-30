package recursive

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
	in recursion we break the problem into 1 and (n-1) parts where we handle 1 part and let recursion handle (n-1) parts,
	so we pass head->next (i.e. from the second node to the end of the list) to the recursive call.
**/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseList(head.Next)
	node.Next.Next = head
	head.Next = nil
	return node
}
