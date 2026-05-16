/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var ans int

func traverse(node *TreeNode) int {

    if node == nil {
        return 0
    }

    leftPath := traverse(node.Left)
    rightPath := traverse(node.Right)

    ans = max(ans, node.Val, node.Val+leftPath, node.Val+rightPath, node.Val+rightPath+leftPath)

    return max(node.Val, node.Val+leftPath, node.Val+rightPath)
    
}

func maxPathSum(root *TreeNode) int {
    ans = math.MinInt

    traverse(root)

    return ans


}
