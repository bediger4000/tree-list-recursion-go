package main

import (
	"fmt"
	"strconv"
	"os"
)

type TreeNode struct {
	data int
	left *TreeNode
	right *TreeNode
}

func insert(node *TreeNode, value int) *TreeNode {

	if node == nil {
		node = new(TreeNode)
		node.data = value
	} else {
		if value < node.data {
			node.left = insert(node.left, value)
		} else {
			node.right = insert(node.right, value)
		}
	}

	return node
}

func convertTree(node *TreeNode) *TreeNode {
	var x TreeNode
	return &x
}

func traverseList(node *TreeNode) {
	for p := node; p != nil; p = p.left {
		fmt.Printf("%d\n", p.data)
	}
}

func traverseTree(node *TreeNode) {
	if node == nil { return }

	traverseTree(node.left)
	fmt.Printf("%d\n", node.data)
	traverseTree(node.right)
}

func main() {
	var root *TreeNode

	for _, str := range os.Args[1:] {
		val, err := strconv.Atoi(str)

		if err == nil {
			root = insert(root, val)
		} else {
			fmt.Printf("Problem with %q: %s\n", str, err)
		}
	}

	if root != nil {
		fmt.Printf("Sorted binary tree traversal:\n")
		traverseTree(root)
		head := convertTree(root)
		if head != nil {
			fmt.Printf("\nList traversal:\n")
			traverseList(head)
		}
	}

}
