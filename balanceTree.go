package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	LeftNode  *TreeNode
	Value     int
	RightNode *TreeNode
}

func main() {
	var tree = &TreeNode{nil, 1, nil}
	tree.insert(2)
	tree.insert(3)
	tree.insert(4)
	tree.insert(5)

	printTree(tree)

	checBalance := is_balanced(tree)
	fmt.Println("is tree balanced: ", checBalance)
}

func is_balanced(tree *TreeNode) bool {
	if tree == nil {
		return false
	}

	lheight := getHeight(tree.LeftNode)
	rheight := getHeight(tree.RightNode)

	heightDiff := math.Abs(float64(lheight) - float64(rheight))
	fmt.Println("heightDiff ", heightDiff)

	if heightDiff > 1 {
		return false
	}

	return is_balanced(tree.LeftNode) && is_balanced(tree.RightNode)
}

func getHeight(tree *TreeNode) float64 {
	if tree == nil {
		return 0
	}

	lheight := getHeight(tree.LeftNode)
	rheight := getHeight(tree.RightNode)

	return math.Max(lheight, rheight) + 1
}

func (tree *TreeNode) insert(val int) {
	if tree != nil {
		if tree.LeftNode == nil {
			tree.LeftNode = &TreeNode{nil, val, nil}
		} else if tree.RightNode == nil {
			tree.RightNode = &TreeNode{nil, val, nil}
		} else if tree.LeftNode != nil {
			tree.LeftNode.insert(val)
		} else if tree.RightNode != nil {
			tree.RightNode.insert(val)
		} else {
			tree = &TreeNode{nil, val, nil}
		}
	}
}

func printTree(tree *TreeNode) {
	if tree != nil {
		fmt.Println("value: ", tree.Value)
		fmt.Println("Tree left node:")
		printTree(tree.LeftNode)
		fmt.Println("Tree right node:")
		printTree(tree.RightNode)
	} else {
		fmt.Print("Nil\n")
	}
}
