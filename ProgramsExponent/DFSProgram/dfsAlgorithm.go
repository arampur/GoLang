package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinNode struct {
	Val      int
	Children []*BinNode
}

func main() {

	root := &Node{1, nil, nil}
	root.Left = &Node{2, nil, nil}
	root.Right = &Node{3, nil, nil}
	root.Left.Left = &Node{4, nil, nil}
	root.Left.Right = &Node{5, nil, nil}
	root.Right.Left = &Node{6, nil, nil}
	root.Right.Right = &Node{7, nil, nil}

	dfsAlgo(root)

	// node4 := &BinNode{4, []*BinNode{}}
	// node5 := &BinNode{5, []*BinNode{}}
	// node2 := &BinNode{2, []*BinNode{node4, node5}}
	// node6 := &BinNode{6, []*BinNode{}}
	// node7 := &BinNode{7, []*BinNode{}}
	// node3 := &BinNode{3, []*BinNode{node6, node7}}
	// binroot := &BinNode{1, []*BinNode{node2, node3}}

	// fmt.Println("For non-binary tree dfs:")
	// dfsAlgoNonBinaryTree(binroot)

}

func dfsAlgo(root *Node) {
	stack := []*Node{root}
	visited := make(map[*Node]bool)

	if root == nil {
		return
	}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[node] {
			continue
		}

		visited[node] = true

		fmt.Print(node.Value)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	fmt.Println()
}

// func dfsAlgoNonBinaryTree(root *BinNode) {
// 	if root == nil {
// 		return
// 	}

// 	stack := []*BinNode{root}
// 	visited := make(map[*BinNode]bool)

// 	for len(stack) > 0 {
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1] // Pop the node from the stack

// 		if visited[node] {
// 			continue
// 		}

// 		visited[node] = true // Mark the node as visited

// 		fmt.Print(node.Val) // Process the node

// 		// Push children to the stack
// 		for _, child := range node.Children {
// 			stack = append(stack, child)
// 		}
// 	}
// 	fmt.Println()
// }
