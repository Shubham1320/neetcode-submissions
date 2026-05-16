/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {

    queue := []*TreeNode{}
    var ans [][]int

    if root == nil {
        return ans
    }

    queue = append(queue, root)

    for len(queue) > 0 {

        tempQueue := []*TreeNode{}
        ansList := []int{}

        for len(queue) > 0 {
            node := queue[0]
            ansList = append(ansList, node.Val)

            if node.Left != nil {
                tempQueue = append(tempQueue, node.Left)
            }

            if node.Right != nil {
                tempQueue = append(tempQueue, node.Right)
            }
            queue = queue[1:]
        }

        ans = append(ans, ansList)
        queue = append(queue, tempQueue...)

    }
    
    return ans
}
