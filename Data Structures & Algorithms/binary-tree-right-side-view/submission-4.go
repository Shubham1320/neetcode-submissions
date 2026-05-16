/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {

    ans := []int{}

    queue := []*TreeNode{}

    if root == nil {
        return ans
    }

    queue = append(queue, root)

    for len(queue) > 0 {

        tempQueue := []*TreeNode{}
        var node *TreeNode

        for len(queue) > 0 {
            node = queue[0]

            if node.Left != nil {
                tempQueue = append(tempQueue, node.Left)
            }

            if node.Right != nil {
                tempQueue = append(tempQueue, node.Right)
            }
            queue = queue[1:]
        }

        ans = append(ans, node.Val)
        queue = append(queue, tempQueue...)

    }
    
    return ans
}
