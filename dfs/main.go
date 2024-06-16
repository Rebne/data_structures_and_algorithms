package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func iterative_inorder_dfs(root *Node) {
	stack := []*Node{}
	curr := root

	for curr != nil || len(stack) > 0 {

		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			n := len(stack) - 1
			curr = stack[n]
			fmt.Println(curr.Value)
			stack = stack[:n]
			curr = curr.Right
		}
	}
}

func main() {
	// Creating a sample binary tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5

	root := &Node{
		Value: 1,
		Left: &Node{
			Value: 2,
			Left: &Node{
				Value: 4,
			},
			Right: &Node{
				Value: 5,
			},
		},
		Right: &Node{
			Value: 3,
		},
	}

	iterative_inorder_dfs(root)

}
