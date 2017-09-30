package main

import (
	"fmt"
	"strconv"
	"os"
)

type TreeNode struct {
	data int
	left *TreeNode   // also prev link
	right *TreeNode  // also next link
}

// Build a sorted binary tree
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

// In-order traverse of binary tree
func traverseTree(node *TreeNode) {
	if node == nil { return }

	traverseTree(node.left)
	fmt.Printf("%d\n", node.data)
	traverseTree(node.right)
}

// Function specifed by problem statement.
func treeToList(root *TreeNode) *TreeNode {

	head, tail := convertTree(root, nil, nil)
	tail.right = head
	head.left = tail

	return head
}

func convertTree(node *TreeNode, head *TreeNode, tail *TreeNode) (*TreeNode, *TreeNode) {

	left := node.left
	right := node.right

	if left != nil {
		head, tail = convertTree(left, head, tail)
	}

	if tail != nil {
		tail.right = node
	}
	node.left  = tail
	tail = node

	if head == nil && left == nil && right == nil {
		head = node
	}

	if right != nil {
		head, tail = convertTree(right, head, tail)
	}

	return head, tail
}

func main() {
	var root *TreeNode

	// Build sorted binary tree from representation of numbers
	// on command line.
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

		head := treeToList(root)

		// Run the list forward and back to prove links
		// got put in correctly. I feel like a do-while
		// loop would look cleaner, or a non-circular
		// doubly-linked list would work cleaner.

		fmt.Printf("\nList traversal:\n")
		fmt.Printf("%d\n", head.data)
		for p := head.right; p != head; p = p.right {
			fmt.Printf("%d\n", p.data)
		}

		fmt.Printf("\nReverse list traversal:\n")
		for p := head.left; p != head; p = p.left {
			fmt.Printf("%d\n", p.data)
		}
		fmt.Printf("%d\n", head.data)
	}

}
