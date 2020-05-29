package leetCodeHard

import "fmt"

func reverseKGroup(head *ListNode, k int) *ListNode {
	head = makeLinkNode([]int{1, 2, 3, 4, 5})
	k = 3

	var root *ListNode
	remain := head
	for nil != remain {
		tmp, tmpRemain := splitN(remain, k)
		if root == nil {
			root = tmp
		} else {
			tailHandle(root, func(node *ListNode) {
				node.Next = tmp
			})
		}
		remain = tmpRemain
	}
	printLinkList(root)
	return root
}

func tailHandle(root *ListNode, handle func(node *ListNode)) {
	tmp := root
	for {
		if tmp.Next != nil {
			tmp = tmp.Next
		} else {
			defer handle(tmp)
			return
		}
	}
}

func splitN(head *ListNode, k int) (*ListNode, *ListNode) {
	var tmp *ListNode
	for c := 0; c < k; c++ {
		if head == nil {
			break
		}
		next := head.Next
		if tmp == nil {
			tmp = head
			tmp.Next = nil
		} else {
			head.Next = tmp
			tmp = head
		}
		head = next
	}
	return tmp, head
}

/********/
func makeLinkNode(list []int) *ListNode {
	var root *ListNode
	tail := root
	for _, v := range list {
		node := ListNode{
			Val: v,
		}
		if root == nil {
			root = &node
			tail = root
		} else {
			tail.Next = &node
			tail = &node
		}
	}
	return root
}

func printLinkList(root *ListNode) {
	for root != nil {
		fmt.Printf("%v ", root.Val)
		root = root.Next
	}
	fmt.Println("")
}
