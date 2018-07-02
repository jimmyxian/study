package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func printByLevel(tree *TreeNode, level int) int {
	if tree == nil || level < 0 {
		return 0
	}
	if level == 0 {
		fmt.Print(tree.Data)
		return 1
	}
	return printByLevel(tree.Left, level-1) + printByLevel(tree.Right, level-1)
}

func printTree(root *TreeNode) {
	for i := 0; i > -1; i++ {
		if printByLevel(root, i) == 0 {
			break
		}
		fmt.Println("----")
	}
}

func printTree2(root *TreeNode) {
	q1 := list.New()
	q2 := list.New()

	q1.PushBack(root)

	for q1.Len() != 0 {
		for q1.Len() != 0 {
			tempData := q1.Remove(q1.Front()).(*TreeNode)
			fmt.Print(tempData.Data)

			if tempData.Left != nil {
				q2.PushBack(tempData.Left)
			}
			if tempData.Right != nil {
				q2.PushBack(tempData.Right)
			}
		}
		fmt.Println(" ")
		q1, q2 = q2, q1
	}
}

func main() {

	root := &TreeNode{Data: 0}
	root.Left = &TreeNode{Data: 1}
	root.Right = &TreeNode{Data: 2}
	root.Left.Left = &TreeNode{Data: 3}
	root.Left.Right = &TreeNode{Data: 4}
	root.Right.Left = &TreeNode{Data: 5}
	root.Right.Right = &TreeNode{Data: 6}
	root.Right.Left.Left = &TreeNode{Data: 7}
	root.Right.Left.Right = &TreeNode{Data: 8}

	printTree2(root)
	//fmt.Println("asdf")
}
