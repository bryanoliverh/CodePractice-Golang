package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil || p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
    root1 := &TreeNode{Val: 1}
    root1.Left = &TreeNode{Val: 2}
    root1.Right = &TreeNode{Val: 3}
    root1.Left.Left = &TreeNode{Val: 4}
    root1.Left.Right = &TreeNode{Val: 5}
    root1.Right.Left = &TreeNode{Val: 6}

    root2 := &TreeNode{Val: 1}
    root2.Left = &TreeNode{Val: 2}
    root2.Right = &TreeNode{Val: 3}
    root2.Left.Left = &TreeNode{Val: 4}
    root2.Left.Right = &TreeNode{Val: 5}
    root2.Right.Left = &TreeNode{Val: 6}

    if isSameTree(root1, root2) {
        fmt.Println("The two binary trees are identical.")
    } else {
        fmt.Println("The two binary trees are different.")
    }

    fmt.Println("Tree 1:")
    printTree(root1, "")
    fmt.Println("Tree 2:")
    printTree(root2, "")
}

func printTree(root *TreeNode, prefix string) {
    if root == nil {
        return
    }
    fmt.Println(prefix, root.Val)
    printTree(root.Left, prefix + "│  ")
    printTree(root.Right, prefix + "   ")
}
