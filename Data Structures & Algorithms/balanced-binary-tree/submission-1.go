/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
} 

//            1

//        2       3

//            4

//         5   

func traverse(node *TreeNode) (int, bool) {

    if node == nil {
        return 0, true
    }

    if node.Left == nil && node.Right == nil {
        return 1, true
    }

    leftDepth, leftAns := traverse(node.Left)
    rightDepth, rightAns := traverse(node.Right)

    fmt.Println(node.Val, leftDepth, leftAns, rightDepth, rightAns)

    if Abs(leftDepth - rightDepth) <= 1 {
        return max(leftDepth, rightDepth)+1, leftAns && rightAns
    }

    return 0, false
}

func isBalanced(root *TreeNode) bool {

    if root == nil {
        return true
    }

    if root.Left == nil && root.Right == nil {
        return true
    }
    _, ans := traverse(root)
    return ans
}
