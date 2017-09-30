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

func convertTree(node *TreeNode, tail *TreeNode) *TreeNode {

	left := node.left
	right := node.right

	if left != nil {
		tail = convertTree(left, tail)
	}

	if tail != nil {
		tail.right = node
	}
	node.left  = tail
	tail = node

	if right != nil {
		tail = convertTree(right, tail)
	}

	return tail
}

func traverseList(node *TreeNode) {
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

		tail := convertTree(root, nil)
		if tail != nil {
			fmt.Printf("\nReverse list traversal:\n")
			for p := tail; p != nil; p = p.left {
				fmt.Printf("%d\n", p.data)
			}

			var head *TreeNode
			for head = tail; head.left != nil; head = head.left { }

			fmt.Printf("\nList traversal:\n")
			for p := head; p != nil; p = p.right {
				fmt.Printf("%d\n", p.data)
			}
		}
	}

}
