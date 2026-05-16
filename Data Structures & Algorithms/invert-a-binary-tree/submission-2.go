/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(node *TreeNode) {

    if node == nil {
        return
    }

    if node.Left != nil && node.Right != nil {
        traverse(node.Left)
        traverse(node.Right)
        
        temp := node.Left
        node.Left = node.Right
        node.Right = temp
    } else if node.Left != nil {
        traverse(node.Left)
        node.Right = node.Left
        node.Left = nil
    } else if node.Right != nil {
        traverse(node.Right)
        node.Left = node.Right
        node.Right = nil
    }

}


func invertTree(root *TreeNode) *TreeNode {

    traverse(root)
    return root
}
