/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(node *TreeNode, left int, right int) bool {

    if node == nil {
        return true
    }

    if !(node.Val < right && node.Val > left) {
        return false
    }

    return traverse(node.Left, left, node.Val) && traverse(node.Right, node.Val, right)
}

func isValidBST(root *TreeNode) bool {

    return traverse(root, math.MinInt, math.MaxInt)

}
