package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func inorder_dfs(root *Node) {
	if root == nil {
		return
	}
	inorder_dfs(root.Left)
	fmt.Println(root.Value)
	inorder_dfs(root.Right)
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

func preorder_dfs(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	preorder_dfs(root.Left)
	preorder_dfs(root.Right)
}

func iterative_preorder_dfs(root *Node) {
	stack := []*Node{}
	curr := root

	for curr != nil || len(stack) > 0 {

		if curr != nil {
			fmt.Println(curr.Value)
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			curr = curr.Left
		} else {
			n := len(stack) - 1
			curr = stack[n]
			stack = stack[:n]
		}
	}
}

func postorder_dfs(root *Node) {
	if root == nil {
		return
	}
	postorder_dfs(root.Left)
	postorder_dfs(root.Right)
	fmt.Println(root.Value)
}
func iterative_postorder_dfs(root *Node) {
	stack := []*Node{root}
	visited := []bool{false}

	for len(stack) > 0 {
		n := len(stack) - 1
		curr, visit := stack[n], visited[n]
		stack, visited = stack[:n], visited[:n]
		if curr != nil {
			if visit {
				fmt.Println(curr.Value)
			} else {
				stack = append(stack, curr)
				visited = append(visited, true)
				stack = append(stack, curr.Right)
				visited = append(visited, false)
				stack = append(stack, curr.Left)
				visited = append(visited, false)
			}
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
	fmt.Println("----------------")
	inorder_dfs(root)
	fmt.Println("----------------")
	iterative_preorder_dfs(root)
	fmt.Println("----------------")
	preorder_dfs(root)
	fmt.Println("----------------")
	iterative_postorder_dfs(root)
	fmt.Println("----------------")
	postorder_dfs(root)
}
