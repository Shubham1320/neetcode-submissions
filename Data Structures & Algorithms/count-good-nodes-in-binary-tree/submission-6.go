/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var ans int

func traverse(node *TreeNode, currMax int) {

    if node == nil {
        return
    }

    fmt.Println(node.Val, currMax)

    if node.Val >= currMax {
        ans++
        currMax = node.Val
    }

    traverse(node.Left, currMax)
    traverse(node.Right, currMax)

}

func goodNodes(root *TreeNode) int {

    ans = 0

    if root == nil {
        return 0
    }

    traverse(root, math.MinInt)

    return ans
}
