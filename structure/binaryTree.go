package structure

type compareResultType int

const (
	compareResultTypeNone compareResultType = iota
	compareResultTypeEqual
	compareResultTypeSmaller
	compareResultTypeBigger
)

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (n *BinaryTreeNode) Compare(ele *BinaryTreeNode) compareResultType {
	if nil == ele {
		return compareResultTypeNone
	}
	if n.value == ele.value {
		return compareResultTypeEqual
	} else if n.value < ele.value {
		return compareResultTypeSmaller
	}
	return compareResultTypeBigger
}
