package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	var root *TreeNode

	root = insertBST(root, 4)
	root = insertBST(root, 7)
	root = insertBST(root, 2)
	root = insertBST(root, 3)
	root = insertBST(root, 1)

	fmt.Println(searchBST(root,2))

}


// search the value is there in the tree, 
// if there then return the root

func searchBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val == val {
		

		return &TreeNode{Val: root.Val, Left: root.Left, Right: root.Right}
	}

	if root.Val < val {

		return searchBST(root.Right,val)
	} else if root.Val > val {

		return searchBST(root.Left, val)
	}

	return nil

}

func displayTree(root *TreeNode) {

	if root != nil {

		displayTree(root.Left)
		fmt.Println(root.Val)
		displayTree(root.Right)
	}
}

func insertBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		root = &TreeNode{Val: val, Left: nil, Right: nil}
	}

	if root.Val < val {
		root.Right = insertBST(root.Right, val)
	} else if root.Val > val {
		root.Left = insertBST(root.Left, val)
	}

	return root
}
