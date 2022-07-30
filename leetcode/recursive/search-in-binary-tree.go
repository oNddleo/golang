package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &TreeNode{Val: data}
	} else {
		t.root.insert(data)
	}
}

func (n *TreeNode) insert(data int) {
	if data <= n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: data}
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Val: data}
		} else {
			n.Right.insert(data)
		}
	}
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > val {
		// search left and right
		return searchBST(root.Left, val)
	}

	if root.Val == val {
		// search and return node value
		return root
	}
	return searchBST(root.Right, val)
}
