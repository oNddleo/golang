package problem

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	result := &ListNode{}
	currentNode := result
	memorized := 0
	for l1 != nil || l2 != nil || memorized > 0 {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		temp := val1 + val2 + memorized
		memorized = temp / 10
		currentNode.Next = &ListNode{Val: temp % 10}
		currentNode = currentNode.Next
	}
	return result.Next
}
