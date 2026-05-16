/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//                4

//        3               5

//    2


var curr int
var ans int

func traverse(node *TreeNode, k int) {

    if node == nil {
        return 
    }

    traverse(node.Left, k)

    curr++    
    if curr == k {
        ans = node.Val
        return
    }

    traverse(node.Right, k)

}


func kthSmallest(root *TreeNode, k int) int {

    curr = 0
    ans = math.MinInt
    traverse(root, k)

    return ans
    

}
