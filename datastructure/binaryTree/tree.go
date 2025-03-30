package main

import (
	"fmt"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func buildTreePreInOrder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	treeNode := &TreeNode{
		Val: preorder[0],
	}
	index := 0
	for inorder[index] != preorder[0] {
		index++
	}

	treeNode.Left = buildTreePreInOrder(preorder[1:index+1], inorder[:index])
	treeNode.Right = buildTreePreInOrder(preorder[index+1:], inorder[index+1:])

	return treeNode
}

func buildTreeInPostOrder(inorder []int, postorder []int) *TreeNode {
	if len(postorder) != len(inorder) {
		return nil
	}

	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	treeNode := &TreeNode{
		Val: postorder[len(postorder)-1],
	}

	index := 0
	for inorder[index] != postorder[len(postorder)-1] {
		index++
	}

	treeNode.Left = buildTreeInPostOrder(inorder[:index], postorder[:index])
	treeNode.Right = buildTreeInPostOrder(inorder[index+1:], postorder[index:len(postorder)-1])

	return treeNode
}

func buildTreePrePostOrder(preorder []int, postorder []int) *TreeNode {
	if len(postorder) != len(preorder) {
		return nil
	}

	if len(postorder) == 0 || len(preorder) == 0 {
		return nil
	}

	// root
	treeNode := &TreeNode{
		Val: preorder[0],
	}

	if len(preorder) == 1 {
		return treeNode
	}

	index := 0
	for postorder[index] != preorder[1] {
		index++
	}

	treeNode.Left = buildTreePrePostOrder(preorder[1:index+2], postorder[:index+1])
	treeNode.Right = buildTreePrePostOrder(preorder[index+2:], postorder[index+1:len(postorder)-1])

	return treeNode
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := buildTreePreInOrder(preorder, inorder)
	fmt.Println(tree)
}
