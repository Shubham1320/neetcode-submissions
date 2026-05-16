/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(p *TreeNode, q *TreeNode) bool {

    if p != nil && q != nil {

        leftAns := traverse(p.Left, q.Left)
        rightAns := traverse(p.Right, q.Right)

        return p.Val == q.Val && leftAns && rightAns
    } else if p == nil && q == nil{
        return true
    } else {
        return false
    }

}

func isSameTree(p *TreeNode, q *TreeNode) bool {

    return traverse(p, q)

}
