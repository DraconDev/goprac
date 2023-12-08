package tree

import (
	"fmt"
)

// type TreeNode struct {
// 	Value int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// Sample binary tree
//        1
//       / \
//      2   3
//     / \ / \
//    4  5 6  7

func BuildSampleTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
}

func BfsBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func dfsBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	var dfsHelper func(node *TreeNode)
	dfsHelper = func(node *TreeNode) {
		if node != nil {
			result = append(result, node.Val)
			dfsHelper(node.Left)
			dfsHelper(node.Right)
		}
	}

	dfsHelper(root)
	return result
}

func main() {
	tree := BuildSampleTree()

	// Example of BFS traversal on a binary tree
	fmt.Println("BFS Traversal:")
	fmt.Println(BfsBinaryTree(tree))

	// Example of DFS traversal on a binary tree
	fmt.Println("DFS Traversal:")
	fmt.Println(dfsBinaryTree(tree))
}
