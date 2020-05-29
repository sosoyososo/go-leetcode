package leetCodeHard

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root TreeNode) [][]int {
	ret := [][]int{}
	list := []TreeNode{root}

	for {
		if len(list) == 0 {
			break
		}
		tmp := []int{}
		tmpNodes := []TreeNode{}
		for _, node := range list {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				left := node.Left
				tmpNodes = append(tmpNodes, *left)
			}
			if node.Right != nil {
				right := node.Right
				tmpNodes = append(tmpNodes, *right)
			}
		}
		list = tmpNodes
		ret = append(ret, tmp)
	}

	retList := [][]int{}
	for i, _ := range ret {
		retList = append(retList, ret[len(ret)-i-1])
	}

	return retList
}
