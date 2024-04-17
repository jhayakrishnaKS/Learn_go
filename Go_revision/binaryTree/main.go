package main

import (
    "fmt"
)

type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

func (node *TreeNode) Insert(value int) *TreeNode {
    if node == nil {
        return &TreeNode{Value: value}
    }

    if value < node.Value {
        node.Left = node.Left.Insert(value)
    } else {
        node.Right = node.Right.Insert(value)
    }
    return node
}

func (node *TreeNode) Search(value int) bool {
    if node == nil {
        return false
    }

    if value == node.Value {
        return true
    } else if value < node.Value {
        return node.Left.Search(value)
    } else {
        return node.Right.Search(value)
    }
}

func (node *TreeNode) InOrderTraversal() {
    if node == nil {
        return
    }

    node.Left.InOrderTraversal()
    fmt.Printf("%d ", node.Value)
    node.Right.InOrderTraversal()
}

func main() {
    root := &TreeNode{Value: 10}
    root.Insert(5)
    root.Insert(15)
    root.Insert(3)
    root.Insert(8)
    root.Insert(12)
    root.Insert(18)

    fmt.Println("In-order traversal:")
    root.InOrderTraversal()
    fmt.Println()

    fmt.Println("Search for value 8:", root.Search(8))
    fmt.Println("Search for value 20:", root.Search(20))
}
