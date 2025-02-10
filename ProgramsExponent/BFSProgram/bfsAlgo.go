package main

import (
	"fmt"
	"strings"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func bfs(root *Node) *Node {
	if root == nil { //to handle edge case
		return root
	}

	//Using hashmap
	hashMap := make(map[int][]*Node)
	hashMap[0] = []*Node{root}

	visited := []*Node{}
	i := 0

	for {
		nodeArr, ok := hashMap[i]
		if !ok {
			break
		}

		for _, node := range nodeArr {
			visited = append(visited, node)
			if node.left != nil {
				_, ok := hashMap[i+1]
				if !ok {
					hashMap[i+1] = []*Node{}
				}
				hashMap[i+1] = append(hashMap[i+1], node.left)
				hashMap[i+1] = append(hashMap[i+1], node.right)
			}
		}
		i += 1
	}

	var result []string
	for _, v := range visited {
		result = append(result, fmt.Sprintf("%d", v.val))
	}
	fmt.Println(strings.Join(result, " "))

	return root
}

func main() {
	root := &Node{
		val: 1,
		left: &Node{
			val: 2,
			left: &Node{
				val: 4,
			},
			right: &Node{
				val: 5,
			},
		},
		right: &Node{
			val: 3,
			left: &Node{
				val: 6,
			},
			right: &Node{
				val: 7,
			},
		},
	}

	bfs(root)
}
