package leetCodeHard

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var root *ListNode
	var rootTail *ListNode
	for {
		if nil == l1 {
			rootTail.Next = l2
			break
		} else if nil == l2 {
			rootTail.Next = l1
			break
		} else {
			if l1.Val < l2.Val {
				if root == nil {
					root = l1
					rootTail = root
				} else {
					rootTail.Next = l1
					rootTail = rootTail.Next
				}
				l1 = l1.Next
			} else {
				if root == nil {
					root = l2
					rootTail = root
				} else {
					rootTail.Next = l2
					rootTail = rootTail.Next
				}
				l2 = l2.Next
			}
		}
	}
	return root
}
