/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
    "slices"
)

func buildTree(preorder []int, inorder []int) *TreeNode {

    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    
    newNode := &TreeNode {
        Val: preorder[0],
    }

    inorderIndex := slices.Index(inorder, preorder[0])
    fmt.Println(inorderIndex)

    newNode.Left = buildTree(preorder[1:inorderIndex+1], inorder[:inorderIndex])
    newNode.Right = buildTree(preorder[inorderIndex+1:], inorder[inorderIndex+1:])

    return newNode

}
