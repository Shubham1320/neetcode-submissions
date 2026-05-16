/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(node *TreeNode) int {

    if node == nil {
        return 0
    }
    
    if node.Left == nil && node.Right == nil {
        return 1
    }

    leftDepth := traverse(node.Left)
    rightDepth := traverse(node.Right)

    fmt.Println(leftDepth,rightDepth)

    return max(leftDepth, rightDepth) + 1

}


func maxDepth(root *TreeNode) int {

    return traverse(root)

}
