/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var diameter = 0

func traverse(node *TreeNode) int {

    if node == nil {
        return 0
    }

    if node.Left == nil && node.Right == nil {
        return 1
    }

    leftDepth := traverse(node.Left)
    rightDepth := traverse(node.Right)

    if leftDepth + rightDepth + 1 > diameter {
        diameter = leftDepth + rightDepth + 1
    }

    fmt.Println(diameter)

    return max(leftDepth, rightDepth) + 1

}

func diameterOfBinaryTree(root *TreeNode) int {
    diameter = 0
    traverse(root)
    if diameter == 0 {
        return diameter
    }
    return diameter - 1
}
