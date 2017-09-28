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

func traverse(node *TreeNode) {
	if node == nil { return }

	traverse(node.left)
	fmt.Printf("%d\n", node.data)
	traverse(node.right)
}

func main() {
	var root *TreeNode

	for _, str := range os.Args[1:] {
		val, err := strconv.Atoi(str)

		if err == nil {
			fmt.Printf("insert %d\n", val)
			root = insert(root, val)
		} else {
			fmt.Printf("Problem with %q: %s\n", str, err)
		}
	}

	if root != nil {
		traverse(root)
	}
}
