package leetcode

import (
    "container/list"
    "fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func ArrayToTree(arr []int) *TreeNode {
    if len(arr) == 0 {
        return nil
    }

    root := &TreeNode{Val: arr[0]}
    queue := list.New()
    queue.PushBack(root)
    i := 1

    for i < len(arr) {
        current := queue.Front().Value.(*TreeNode)
        queue.Remove(queue.Front())

        leftVal := arr[i]
        i++
        if leftVal != -1 { // -1 can be used to represent a null node
            current.Left = &TreeNode{Val: leftVal}
            queue.PushBack(current.Left)
        }

        if i < len(arr) {
            rightVal := arr[i]
            i++
            if rightVal != -1 { // -1 can be used to represent a null node
                current.Right = &TreeNode{Val: rightVal}
                queue.PushBack(current.Right)
            }
        }
    }

    return root
}

func LevelOrderTraversal(root *TreeNode) {
    if root == nil {
        return
    }

    queue := list.New()
    queue.PushBack(root)

    for queue.Len() > 0 {
        current := queue.Front().Value.(*TreeNode)
        fmt.Printf("%d ", current.Val)
        queue.Remove(queue.Front())

        if current.Left != nil {
            queue.PushBack(current.Left)
        }
        if current.Right != nil {
            queue.PushBack(current.Right)
        }
    }
}

func PrintTree(root *TreeNode, indent string, last bool) {
    if root == nil {
        return
    }

    fmt.Print(indent)

    if last {
        fmt.Printf("└── ")
        indent += "    "
    } else {
        fmt.Printf("├── ")
        indent += "│   "
    }

    fmt.Println(root.Val)

    PrintTree(root.Left, indent, false)
    PrintTree(root.Right, indent, true)
}

